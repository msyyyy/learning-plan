双向链表+哈希表 



新数据或击中的数据放到链表头部，表示最近使用的数据，如果链表满，从尾部淘汰数据。但只用链表会存在一个问题，击中数据的时间复杂度为O(n)，每次需要遍历链表，所以引入哈希表，时间复杂度降到O(1)，以空间换时间。



```c++
#include <list>
#include <unordered_map>
#include <cassert>   
using namespace std;  
struct Element
{
    int key;
    int value;
    Element(int k, int v):key(k), value(v){}
};
class LRUCache {
private:
    list<Element> m_list;
    unordered_map<int, list<Element>::iterator> m_map;
    int m_capacity;
public:
    LRUCache(int capacity) {
        m_capacity = capacity;
    }

    int get(int key) {
        if (m_map.find(key) == m_map.end())
            return -1;
        else
        {
            //将元素放入链表头部
            m_list.splice(m_list.begin(), m_list, m_map[key]);
            m_map[key] = m_list.begin();
            return m_list.begin()->value;
        }
    }

    void put(int key, int value) {
        assert(m_capacity > 0);
        if (m_map.find(key) != m_map.end())
        {   //更value
            m_map[key]->value = value;
            //将元素放入链表头部
            m_list.splice(m_list.begin(), m_list, m_map[key]);
            m_map[key] = m_list.begin();
        }
        else if (m_capacity == m_list.size())
        {
            m_map.erase(m_list.back().key);
            m_list.pop_back();
            m_list.push_front(Element(key, value));
            m_map[key] = m_list.begin();
        }
        else
        {
            m_list.push_front(Element(key, value));
            m_map[key] = m_list.begin();
        }
    }
};

```

