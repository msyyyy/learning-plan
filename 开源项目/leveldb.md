## 整体把握

[整体把握](<https://www.jianshu.com/p/e904834932a1>)

[介绍](<https://blog.csdn.net/qq_26499321/article/details/78063856>)

![img](https://upload-images.jianshu.io/upload_images/299848-58547f271e92bd87.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/851/format/webp)



- 写

write()

①.会先写到 log文件中 ，log文件在磁盘上，所以不怕数据消失 ，②.然后将数据写入memtable （一个叫跳跃表的数据结构） ,memtable 在内存上，然后write操作成功

当③.memtable大小到达一个上限时，memtable会变成  Immutale memtable  ，Immutale memtable时只读的，会重新创建一个新的memtable  ，④Immutale memtable会定期刷到磁盘，然后清除对应的log文件，产生新的log文件 ⑤磁盘中数据存储格式为sstable

**sstable是怎么样的**

![img](https://upload-images.jianshu.io/upload_images/299848-0caeda1bfac8a28b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1000/format/webp)

图中红色框部分就是他的样子。

level0是由Imuabletable memtable直接dump到磁盘中的，level1是由level0经过compaction获得，level2是由level1经过compaction获得，以此类推。其中每个文件后缀为`.sst`, 每个level中的文件数都是有限制的，超过了限制则会被compaction到更高level的层次上去，所以这个东西叫leveldb。其中每一个level符合以下规则：

```
1. level0中的单个文件（sst）是有序的，但是文件与文件之间是无序的并且有可能有重合的key
2. level 1 ~ level n 每一个level中在自己level中都是全局有序的
3. mainifest文件中包含了每一个sst文件的最小key和最大key，方便查找
```

- 读

read()

```
1. memtable
2. Imuable memtable
3. level 0
4. level 1 
 ......
```

上面步骤非常清晰。首先查找会先经过memtable查找，找不到就依次按顺序往下找，一直找不到就返回empty了。当然，中间有很多细节优化，这里先不深入去理解，例如会通过布隆算法过滤掉不存在的key。后面讲到filter的时候会深入讲解。

- 删除

  > leveldb没有实际的删除操作，就是write一个删除标记和key进去

  啥意思？我们都知道所有key的插入都是时间有序的，从memtable到level n一路飞奔，从不回头。
   所以我们假如我们添加了key为`name`, value 为`戈风`的数据进去之后，我们要删除key为`name`的记录，我们只需要再插入一条 `name del（标志）`，这样我们查找的时候就会先遇到`name del`表示key已经被删除，返回empty。

  除了这个时间有序之外，在level n compaction到level n + 1的时候如果发现有删除的key，这时候就会真正删除它。

**重点模块**

```
1. memtable
2. Imuable memtable（其实我跟1是亲兄弟啦）
3. sstable
4. log
5. filter
```

- db/, 数据库逻辑
- doc/, MD文档
- helpers/, LevelDB内存版, 通过namespace覆盖
- port/, 平台相关代码
- table/, LSM有关的

- The basic operations are `Put(key,value)`, `Get(key)`, `Delete(key)`.

## 使用

### 头文件准备

① 找到静态链接库  `libleveldb.a`

② 静态编译如下 ，静态链接库必须在当前目录下，或者指出库的绝对路径

最好指定c++11编译，链接到`-lpthread`，因为`leveldb`有用到线程相关调用。

`g++ -g -Wall -std=c++11 dd.cpp -o mytest ./libleveldb.a -lpthread`

### LSM TREE

**LSM树和B+树相比，LSM树牺牲了部分读性能，用来大幅提高写性能。**

LSM树的设计思想非常朴素：**将对数据的修改增量保持在内存中，达到指定的大小限制后将这些修改操作批量写入磁盘**

所有数据直接写入memtable并暂时持久化到磁盘(log), 当memtable足够大的时候, 变为immemtable, 开始往硬盘挪, 成为SSTable. 这就是LSM Tree仅有的全部. 你可以用任何有道理的数据结构来表示memtable, immemtable和SSTable. Google选择用跳表实现memtable和immemtable, 用有序行组来实现SSTable.



1. 添加BloomFilter, 这样可以提升全库扫描的速度, 肯定没这个key的SSTable直接跳过.

2. leveled compaction, 把SSTable分成不同的等级. 除等级0以外, 其余各等级的SSTable不会有重复的key.

```
#include "leveldb/db.h"
#include <cassert>
#include <iostream>

using namespace std;
using namespace leveldb;

int main() {
    leveldb::DB *db;
    leveldb::Options options;
    options.create_if_missing = true;
    leveldb::Status status = leveldb::DB::Open(options, "testdb", &db);
    assert(status.ok());

    status = db->Put(WriteOptions(), "JimZuoLin", "Hello Jim!");
    assert(status.ok());
    string res;
    status = db->Get(ReadOptions(), "JimZuoLin", &res);
    assert(status.ok());
    cout << res << endl;

    delete db;
    return 0;
}
```

- 内存池 Arena  ，在 `uti/arena.cc中   `主要就是减少malloc或者new调用的次数，较少内存分配所带来的系统开销。

  类的成员变量

  ```c++
  char* alloc_ptr_;//内存的偏移量指针，即指向未使用内存的首地址
  size_t alloc_bytes_remaining_;//还剩下的内存数
  std::vector<char*> blocks_;//存储每一次分配的内存指针
  size_t blocks_memory_;//到目前为止分配的总内存。
  ```

  接下来分析下Arena内存分配的主要函数。

  ```c++
  public:
  	char* Allocate(size_t bytes);
  private:
  	char* AllocateFallback(size_t bytes);
  	char* AllocateNewBlock(size_t block_bytes);
  ```

  Arena对外提供的接口是public里的函数，但是该函数会调用private里的两个函数.我分析内存分配策略。当要分配内存的时候：

  1. 如果需求的内存小于剩余的内存，那么直接在剩余的内存分配就可以了；
  2. 如果需求的内存大于剩余的内存，而且大于4096/4，则给这内存单独分配一块bytes（函数参数）大小的内存。
  3. 如果需求的内存大于剩余的内存，而且小于4096/4，则重新分配一个内存块，默认大小4096，用于存储数据。

  针对第二点，按源码的注释是说避浪费太多的剩余空间。我的理解是，如果剩余的内存为1500kb，那么假设有一个内存需求是500kb，一个内存需求是1500kb,则第一个需求可以使用三次才导致进行一次重新内存分配，而第二个只能使用一次就要进行一次重新内存分配。所以leveldb第二条的用意主要还是减少内存分配的次数。

- leveldb在内存中存储数据的区域称为memtable，这个memtable底层是用跳跃链表skiplist来实现的。

  leveldb是支持多线程操作的，但是skiplist并没有使用linux下锁，信号量来实现同步控制，据说是因为锁机制导致某个线程占有资源，其他线程阻塞的情况，导致系统资源利用率降低。所以leveldb采用的是**内存屏障**来实现同步机制

  ## skiplist节点

  接下来，我们看下skiplist节点源代码：

  ```c++
  template<typename Key, class Comparator>
  struct SkipList<Key,Comparator>::Node {
    explicit Node(const Key& k) : key(k) { }
  
    Key const key;
  
    // Accessors/mutators for links.  Wrapped in methods so we can
    // add the appropriate barriers as necessary.
    Node* Next(int n) {
      assert(n >= 0);
      // Use an 'acquire load' so that we observe a fully initialized
      // version of the returned Node.
      return reinterpret_cast<Node*>(next_[n].Acquire_Load());
    }//获取当前节点的下一个节点
  
    void SetNext(int n, Node* x) {
      assert(n >= 0);
      // Use a 'release store' so that anybody who reads through this
      // pointer observes a fully initialized version of the inserted node.
      next_[n].Release_Store(x);//设置当前节点的下个节点。
    }
  
    // No-barrier variants that can be safely used in a few locations.
    Node* NoBarrier_Next(int n) {
      assert(n >= 0);
      return reinterpret_cast<Node*>(next_[n].NoBarrier_Load());
    }//无需内存屏障的查找下一个节点
  
    void NoBarrier_SetNext(int n, Node* x) {
      assert(n >= 0);
      next_[n].NoBarrier_Store(x);
    }无需内存屏障的设置下一个节点
  
   private:
    // Array of length equal to the node height.  next_[0] is lowest level link.
    port::AtomicPointer next_[1];//节点的层数。
  };
  ```

  

  这是skiplist节点类，可以查找下一个节点，可以设置下一个节点。nect_[1]是一个很优秀的设计。只定义数组第一个节点，然后分配内存时，分配（高度-1）个数组类型的内存。其实就是动态分配内存。否则一开始用数组分配大数组，易造成内存浪费。

  ## skiplist实现

  skiplist成员变量有：

  ```c++
    enum { kMaxHeight = 12 };//定义skiplist链表最高节点
  // Immutable after construction
  Comparator const compare_;//比较器有最顶层的options通过一层一层传递下来，用于///链表排序
  Arena* const arena_;    // leveldb内存池，从memtable传过来
  
  Node* const head_;//skiplist头节点
  
  // Modified only by Insert().  Read racily by readers, but stale
  // values are ok.
  port::AtomicPointer max_height_;   // skiplist目前的最高高度
    // Read/written only by Insert().
  Random rnd_;//随机类，用于随机化一个节点高度
  ```

  

  `enum{kMaxHeight = 12}`也是很优秀设计，《efficetive C++》有一个章节有讲到尽量不要使用宏定义定义常量，取代的办法是const常量和enum类型。

  skiplist构造函数如下;

  ```c++
  template<typename Key, class Comparator>
  SkipList<Key,Comparator>::SkipList(Comparator cmp, Arena* arena)
      : compare_(cmp),
        arena_(arena),
        head_(NewNode(0 /* any key will do */, kMaxHeight)),
        max_height_(reinterpret_cast<void*>(1)),
        rnd_(0xdeadbeef) {
    for (int i = 0; i < kMaxHeight; i++) {
      head_->SetNext(i, NULL);
    }
  }
  ```

  

  cmp和arena_都是调用者传进来，head_头指针key初始化为0,高度为链表高度上限。max_height_初始化为1.for循环将头节点的前一个节点都设为NULL。

  分配一个新节点的源码如下：

  ```c++
  template<typename Key, class Comparator>
  typename SkipList<Key,Comparator>::Node*
  SkipList<Key,Comparator>::NewNode(const Key& key, int height) {
    char* mem = arena_->AllocateAligned(
        sizeof(Node) + sizeof(port::AtomicPointer) * (height - 1));//从内存池里面分配
  //足够的内存，用于存储新节点。
    return new (mem) Node(key);//返回这个节点。
  }
  ```

  

  接来看下skiplist插入一个新节点的代码，思想是，在插入高度为height的节点时，首先要找到这个节点height个前节点，然后插入就和普通的链表插入一样。

  ```c++
  template<typename Key, class Comparator>
  void SkipList<Key,Comparator>::Insert(const Key& key) {
    Node* prev[kMaxHeight];//kMaxHeight个前节点，因为高度还未知，所以先设为最大值
    Node* x = FindGreaterOrEqual(key, prev);//查找key值节点前GetMaxHeight()个前节点。
  
    assert(x == NULL || !Equal(key, x->key));
  
    int height = RandomHeight();//随机化一个节点高度
    if (height > GetMaxHeight()) {//如果当前节点的高度大于最高节点，则高出部分的的前节
  //点都是头节点。
      for (int i = GetMaxHeight(); i < height; i++) {
        prev[i] = head_;
      }
      max_height_.NoBarrier_Store(reinterpret_cast<void*>(height));
    }
  
    x = NewNode(key, height);//新建节点
    for (int i = 0; i < height; i++) {
      x->NoBarrier_SetNext(i, prev[i]->NoBarrier_Next(i))//设立当前节点的后节点;
      prev[i]->SetNext(i, x);//设立当前节点的前节点。
    }
  }
  ```

  

  插入节点逻辑其实和单链表类似，只是查找前后节点麻烦点，需要找到不止一个。接下来是判断链表是否含有某个key值的接口：

  ```c++
  template<typename Key, class Comparator>
  bool SkipList<Key,Comparator>::Contains(const Key& key) const {
    Node* x = FindGreaterOrEqual(key, NULL);//返回第0层上个节点
    if (x != NULL && Equal(key, x->key)) {//如果x不为NULL,且key值与x的key值相等，则说明key在链表中。
      return true;
    } else {
      return false;
    }
  }
  ```

  

  链表是没有删除接口的，但是有删除功能。因为当我们插入数据时，key的形式为key:value，当删除数据时，则插入key:deleted类似删除的标记，等到Compaction再删除。

  ## skiplist迭代器

  leveldb用的最多的就是迭代器了，最后我们来分析链表迭代器，分析迭代器，为看STL源码也积累基础。声明如下：

  ```c++
  class Iterator {
     public:
      // Initialize an iterator over the specified list.
      // The returned iterator is not valid.
      explicit Iterator(const SkipList* list);
  
      // Returns true iff the iterator is positioned at a valid node.
      bool Valid() const;
  
      // Returns the key at the current position.
      // REQUIRES: Valid()
      const Key& key() const;
  
      // Advances to the next position.
      // REQUIRES: Valid()
      void Next();
  
      // Advances to the previous position.
      // REQUIRES: Valid()
      void Prev();
  
      // Advance to the first entry with a key >= target
      void Seek(const Key& target);
  
      // Position at the first entry in list.
      // Final state of iterator is Valid() iff list is not empty.
      void SeekToFirst();
  
      // Position at the last entry in list.
      // Final state of iterator is Valid() iff list is not empty.
      void SeekToLast();
  
     private:
      const SkipList* list_;//迭代器迭代的跳跃链表
      Node* node_;//指向当前的节点
      // Intentionally copyable
    };
  ```

  

  迭代器的成员变量有，需要遍历的链表，以及指向当前节点的节点指针。

  链表迭代器构造函数如下

  ```c++
  template<typename Key, class Comparator>
  inline SkipList<Key,Comparator>::Iterator::Iterator(const SkipList* list) {
    list_ = list;
    node_ = NULL;
  }
  ```

  

  迭代器具体实现代码如下：

  ```c++
  template<typename Key, class Comparator>
  inline bool SkipList<Key,Comparator>::Iterator::Valid() const {
    return node_ != NULL;
  }//判断迭代器当前节点是否有效
  
  template<typename Key, class Comparator>
  inline const Key& SkipList<Key,Comparator>::Iterator::key() const {
    assert(Valid());
    return node_->key;
  }//返回当前节点的key值
  
  template<typename Key, class Comparator>
  inline void SkipList<Key,Comparator>::Iterator::Next() {
    assert(Valid());
    node_ = node_->Next(0);
  }//跳跃链表的第0层就是单链表，所以可以直接指向下一个节点
  
  template<typename Key, class Comparator>
  inline void SkipList<Key,Comparator>::Iterator::Prev() {
    // Instead of using explicit "prev" links, we just search for the
    // last node that falls before key.
    assert(Valid());
    node_ = list_->FindLessThan(node_->key);
    if (node_ == list_->head_) {
      node_ = NULL;
    }
  }//查找当前节点的上一个节点。
  
  template<typename Key, class Comparator>
  inline void SkipList<Key,Comparator>::Iterator::Seek(const Key& target) {
    node_ = list_->FindGreaterOrEqual(target, NULL);
  }//查找某个特定的key值的节点。
  
  template<typename Key, class Comparator>
  inline void SkipList<Key,Comparator>::Iterator::SeekToFirst() {
    node_ = list_->head_->Next(0);
  }//查找第一个节点
  
  template<typename Key, class Comparator>
  inline void SkipList<Key,Comparator>::Iterator::SeekToLast() {
    node_ = list_->FindLast();
    if (node_ == list_->head_) {
      node_ = NULL;
    }
  }//最后一个节点
  ```

[CMake](<https://www.jianshu.com/p/aaa19816f7ad>)

1. 查看`leveldb::Status status = leveldb::DB::Open(options, "testdb", &db);`会触发什么模块.

   `leveldb::DB::Open` 在 `db`文件夹 下`db_impl.cc文件中`1469行

   ```c++
   Status DB::Open(const Options& options, const std::string& dbname, DB** dbptr) { // 是工厂函数，创建了对象
     *dbptr = nullptr;
   
     DBImpl* impl = new DBImpl(options, dbname);
     impl->mutex_.Lock();
     VersionEdit edit;
     // Recover handles create_if_missing, error_if_exists
     bool save_manifest = false;
     Status s = impl->Recover(&edit, &save_manifest);
     if (s.ok() && impl->mem_ == nullptr) {
       // Create new log and a corresponding memtable.
       uint64_t new_log_number = impl->versions_->NewFileNumber();
       WritableFile* lfile;
       s = options.env->NewWritableFile(LogFileName(dbname, new_log_number),
                                        &lfile);
       if (s.ok()) {
         edit.SetLogNumber(new_log_number);
         impl->logfile_ = lfile;
         impl->logfile_number_ = new_log_number;
         impl->log_ = new log::Writer(lfile);
         impl->mem_ = new MemTable(impl->internal_comparator_);
         impl->mem_->Ref();
       }
     }
     if (s.ok() && save_manifest) {
       edit.SetPrevLogNumber(0);  // No older logs needed after recovery.
       edit.SetLogNumber(impl->logfile_number_);
       s = impl->versions_->LogAndApply(&edit, &impl->mutex_);
     }
     if (s.ok()) {
       impl->DeleteObsoleteFiles();
       impl->MaybeScheduleCompaction();
     }
     impl->mutex_.Unlock();
     if (s.ok()) {
       assert(impl->mem_ != nullptr);
       *dbptr = impl;
     } else {
       delete impl;
     }
     return s;
   }
   
   Snapshot::~Snapshot() = default;
   ```

   `new DBImpl`的构造函数在 `db`文件夹 下`db_impl.cc文件中`127行

   ```c++
   DBImpl::DBImpl(const Options& raw_options, const std::string& dbname)
       : env_(raw_options.env),
         internal_comparator_(raw_options.comparator),
         internal_filter_policy_(raw_options.filter_policy),
         options_(SanitizeOptions(dbname, &internal_comparator_,
                                  &internal_filter_policy_, raw_options)),
         owns_info_log_(options_.info_log != raw_options.info_log),
         owns_cache_(options_.block_cache != raw_options.block_cache),
         dbname_(dbname),
         table_cache_(new TableCache(dbname_, options_, TableCacheSize(options_))),
         db_lock_(nullptr),
         shutting_down_(false),
         background_work_finished_signal_(&mutex_),
         mem_(nullptr),
         imm_(nullptr),
         has_imm_(false),
         logfile_(nullptr),
         logfile_number_(0),
         log_(nullptr),
         seed_(0),
         tmp_batch_(new WriteBatch),
         background_compaction_scheduled_(false),
         manual_compaction_(nullptr),
         versions_(new VersionSet(dbname_, &options_, table_cache_,
                                  &internal_comparator_)) {}
   
   
   env_, 负责所有IO, 比如建立文件
   internal_comparator_, 用来比较不同key的大小
   internal_filter_policy_, 可自定义BloomFilter
   options_, 将调用者传入的options再用一个函数调整下, 可见Google程序员也不是尽善尽美的... 库的作者要帮忙去除错误参数和优化...
   db_lock_, 文件锁
   shutting_down_, 基于memory barrier的原子指针
   bg_cv_, 多线程的条件
   mem_ = memtable, imm = immemtable
   tmp_batch_, 所有Put都是以batch写入, 这里建立个临时的
   manual_compaction_, 内部开发者调用时的魔法参数, 可以不用理会
   我决定先搞懂memory barrier的原子指针再继续分析
   ```

2. 数据库为了保证运行时崩溃而数据不丢失，会选择写日志，分析leveldb怎么从日志中恢复数据的

   - 人类可读的日志, 存于"LOG"文件

   比如, "2017/06/16-11:09:03.295840 7fffb990d3c0 Recovering log #18". 在代码中是用"Log"函数来触发的, 相关的类是"Logger".

   - 机读二进制日志, 存于".log"文件

   这个是真正意义上用于恢复数据的日志. 数据启动时, 如果有没清空的日志, 就说明上次关闭不成功, 须回放一遍.

   - leveldb::log, 这是一个namespace, 用于把二进制数据安全地序列化, 反序列化

   ::log不仅负责(反)序列化机读日志, VersionEdit在"MANIFEST"文件内也复用了这个组件. 现在提出一个很重要也很常见的问题, 如何保证非原子性的一连串操作的原子性? 有点绕? 来个情景.

   数据库现在要开始写Log了, 一条一条又一条, 这时候突然崩溃了. 下次再开, 日志回放的时候, 会得到啥? 形象的说, 这可以叫做"薛定谔的数据库". 最后一条记录处于成功和失败的叠加态, 只有观测的一瞬间才知道. 大部分用户可以容许的是丢日志, 但绝对不容忍错误的日志被当成正常的写入数据库. 比如, 往A账户转入10000W, 这条写到一半, 最后变成了往A账户转入10W...

   解决方案  checksum

   在数据写入完成之后, 再多写一段hash. 再次读取时, 只有hash和内容对上了, 这段数据才是合法的.

   LevelDB对此有一个高度优化的`crc32c hash`函数在`crc32c.cc`文件内.

   所以, 一条机读日志从内存到硬盘是这样的, 内存对象 => 二进制数组(Slice对象) => leveldb::log切割成小块并打上hash => 写入硬盘.

   实现在`DBImpl::Recover`函数中, 先看`db_impl.cc`280-369行,

   ```c++
   Status DBImpl::Recover(VersionEdit* edit, bool* save_manifest) {
     mutex_.AssertHeld();
   
     // Ignore error from CreateDir since the creation of the DB is
     // committed only when the descriptor is created, and this directory
     // may already exist from a previous failed creation attempt.
     env_->CreateDir(dbname_);// 有可能是第一次打开数据库, 尝试创建目录
     assert(db_lock_ == nullptr);
     Status s = env_->LockFile(LockFileName(dbname_), &db_lock_);
     if (!s.ok()) {
       return s;
     }
   /*
   LevelDB的数据库(文件)是一个文件夹, 如果有多个程序打开这个数据库肯定糟糕了. 那怎么确保一个实例能稳定霸占一个文件夹呢? 跟所有别的程序一样, 建立一个单独的文件, 以独占的方式打开作为LOCK. 每个实例都要尝试创建/拥有这个锁文件, 如果失败了, 说明有别的实例在使用这个数据库.
   
   db_lock_看着很唬人, 其实就是一个简单的结构体, 保存了fd和文件名.
   */
     if (!env_->FileExists(CurrentFileName(dbname_))) {
       if (options_.create_if_missing) {// 没有CURRENT文件, 新建数据库
         s = NewDB(); // 下接NEWDB()
         if (!s.ok()) {
           return s;
         }
       } else {
         return Status::InvalidArgument(
             dbname_, "does not exist (create_if_missing is false)");
       }
     } else {
       if (options_.error_if_exists) {
         return Status::InvalidArgument(dbname_,
                                        "exists (error_if_exists is true)");
       }
     }
   
     s = versions_->Recover(save_manifest);// 开始读之前的manifest
       
     /*
   基于LSM Tree的数据库在恢复时一定分两步, 第一是恢复SSTable, 第二是恢复memtable/immemtable. "versions_->Recover"是前者, 跟入version_set.cc第905行,下接versions_->Recover
     */
     if (!s.ok()) {
       return s;
     }
     SequenceNumber max_sequence(0);
   
     // Recover from all newer log files than the ones named in the
     // descriptor (new log files may have been added by the previous
     // incarnation without registering them in the descriptor).
     //
     // Note that PrevLogNumber() is no longer used, but we pay
     // attention to it in case we are recovering a database
     // produced by an older version of leveldb.
     const uint64_t min_log = versions_->LogNumber();
     const uint64_t prev_log = versions_->PrevLogNumber();
     std::vector<std::string> filenames;
     s = env_->GetChildren(dbname_, &filenames);
     if (!s.ok()) {
       return s;
     }
     std::set<uint64_t> expected;
     versions_->AddLiveFiles(&expected);
     uint64_t number;
     FileType type;
     std::vector<uint64_t> logs;
     for (size_t i = 0; i < filenames.size(); i++) {
       if (ParseFileName(filenames[i], &number, &type)) {
         expected.erase(number);
         if (type == kLogFile && ((number >= min_log) || (number == prev_log)))
           logs.push_back(number);
       }
     }
     if (!expected.empty()) {
       char buf[50];
       snprintf(buf, sizeof(buf), "%d missing files; e.g.",
                static_cast<int>(expected.size()));
       return Status::Corruption(buf, TableFileName(dbname_, *(expected.begin())));
     }
   
     // Recover in the order in which the logs were generated
     std::sort(logs.begin(), logs.end());
     for (size_t i = 0; i < logs.size(); i++) {
       s = RecoverLogFile(logs[i], (i == logs.size() - 1), save_manifest, edit,
                          &max_sequence);
       if (!s.ok()) {
         return s;
       }
   
       // The previous incarnation may not have written any MANIFEST
       // records after allocating this log number.  So we manually
       // update the file number allocation counter in VersionSet.
       versions_->MarkFileNumberUsed(logs[i]);
     }
   
     if (versions_->LastSequence() < max_sequence) {
       versions_->SetLastSequence(max_sequence);
     }
   
     return Status::OK();
   }
   ```

   `NEWDB()`

   ```c++
   Status DBImpl::NewDB() {
     VersionEdit new_db;
     new_db.SetComparatorName(user_comparator()->Name());// 防止用不同comparator打开数据库
     new_db.SetLogNumber(0);
     new_db.SetNextFile(2); // 硬盘文件名后缀, manifest已经占用1了, 所以这里要是2
     new_db.SetLastSequence(0);
   
     const std::string manifest = DescriptorFileName(dbname_, 1);
     WritableFile* file;
     Status s = env_->NewWritableFile(manifest, &file);
     if (!s.ok()) {
       return s;
     }
     {
       log::Writer log(file);
       std::string record;
       new_db.EncodeTo(&record);// 将VersionEdit序列化
       s = log.AddRecord(record);// 写入硬盘
       if (s.ok()) {
         s = file->Close();
       }
     }
     delete file;
     if (s.ok()) {
       // Make "CURRENT" file that points to the new manifest file.
       s = SetCurrentFile(env_, dbname_, 1);// CURRENT文件承担了一个引导的作用
     } else {
       env_->DeleteFile(manifest);
     }
     return s;
   }
   
   //把问题扔给操作系统吧. 锁/引导都用文件系统的原子性和健壮性解决.
   /*
   VersionEdit在LevelDB是什么个概念?
   
   由于LSM Tree没有任何主索引体系, 只要Log+SSTable正确, 就一定能得出正确的结果. 所以不同version之间的差别就是SSTable的差别, A版本到B版本, 可能就是删除名叫e, d, f的SSTable, 再加上o, p, q的SSTable. 这就天然具有了超强的健壮性!
   
   VersionEdit_t0 + VersionEdit_t1 = Data_t1
   
   
   管理如此众多SSTable的任务可以直接甩锅给文件系统, 也就不难理解LevelDB的超强性能了.
   */
   ```

   `version_set.cc`中第883行 

   ```c++
   Status VersionSet::Recover(bool* save_manifest) {
     struct LogReporter : public log::Reader::Reporter { // 作用类似别的语言的委托/回调
       Status* status;
       void Corruption(size_t bytes, const Status& s) override {
         if (this->status->ok()) *this->status = s;
       }
     };
   
     // Read "CURRENT" file, which contains a pointer to the current manifest file
     std::string current;
     Status s = ReadFileToString(env_, CurrentFileName(dbname_), &current);
     if (!s.ok()) {
       return s;
     }
     if (current.empty() || current[current.size() - 1] != '\n') {
       return Status::Corruption("CURRENT file does not end with newline");
     }
     current.resize(current.size() - 1);
   
     std::string dscname = dbname_ + "/" + current;
     SequentialFile* file;
     s = env_->NewSequentialFile(dscname, &file);
     if (!s.ok()) {
       if (s.IsNotFound()) {
         return Status::Corruption("CURRENT points to a non-existent file",
                                   s.ToString());
       }
       return s;
     }
   
     bool have_log_number = false;
     bool have_prev_log_number = false;
     bool have_next_file = false;
     bool have_last_sequence = false;
     uint64_t next_file = 0;
     uint64_t last_sequence = 0;
     uint64_t log_number = 0;
     uint64_t prev_log_number = 0;
     Builder builder(this, current_);
   
     {
       LogReporter reporter;
       reporter.status = &s;
       log::Reader reader(file, &reporter, true /*checksum*/,
                          0 /*initial_offset*/);
       Slice record;
       std::string scratch;
       while (reader.ReadRecord(&record, &scratch) && s.ok()) {
         VersionEdit edit;
         s = edit.DecodeFrom(record);
         if (s.ok()) {
           if (edit.has_comparator_ &&
               edit.comparator_ != icmp_.user_comparator()->Name()) {
             s = Status::InvalidArgument(
                 edit.comparator_ + " does not match existing comparator ",
                 icmp_.user_comparator()->Name());
           }
         }
   
         if (s.ok()) {
           builder.Apply(&edit);
         }
   
         if (edit.has_log_number_) {
           log_number = edit.log_number_;
           have_log_number = true;
         }
   
         if (edit.has_prev_log_number_) {
           prev_log_number = edit.prev_log_number_;
           have_prev_log_number = true;
         }
   
         if (edit.has_next_file_number_) {
           next_file = edit.next_file_number_;
           have_next_file = true;
         }
   
         if (edit.has_last_sequence_) {
           last_sequence = edit.last_sequence_;
           have_last_sequence = true;
         }
       }
     }
     delete file;
     file = nullptr;
   
     if (s.ok()) {
       if (!have_next_file) {
         s = Status::Corruption("no meta-nextfile entry in descriptor");
       } else if (!have_log_number) {
         s = Status::Corruption("no meta-lognumber entry in descriptor");
       } else if (!have_last_sequence) {
         s = Status::Corruption("no last-sequence-number entry in descriptor");
       }
   
       if (!have_prev_log_number) {
         prev_log_number = 0;
       }
   
       MarkFileNumberUsed(prev_log_number);
       MarkFileNumberUsed(log_number);
     }
   
     if (s.ok()) {
       Version* v = new Version(this);
       builder.SaveTo(v);
       // Install recovered version
       Finalize(v);
       AppendVersion(v);
       manifest_file_number_ = next_file;
       next_file_number_ = next_file + 1;
       last_sequence_ = last_sequence;
       log_number_ = log_number;
       prev_log_number_ = prev_log_number;
   
       // See if we can reuse the existing MANIFEST file.
       if (ReuseManifest(dscname, current)) {
         // No need to save new manifest
       } else {
         *save_manifest = true;
       }
     }
   
     return s;
   }
   ```

   