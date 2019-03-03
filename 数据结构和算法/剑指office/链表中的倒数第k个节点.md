##题目描述
输入一个链表，输出该链表中倒数第k个结点。

## 思路

用两个指针a,b

a先前进k-1步，然后a和b同时前进，当a到达尾部时，b所在位置为倒数第k

## 注意

1. k 为 unsigned int 且可能为0

2. k可能超过链表长度

3. 链表头可能为空

```c++
/*
struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(NULL) {
	}
};*/
class Solution {
public:
    ListNode* FindKthToTail(ListNode* pListHead, unsigned int k) {
        if(pListHead==nullptr||k==0) return nullptr;
        ListNode * a=pListHead;
        ListNode * b=pListHead;
        for(unsigned i=0;i<k-1;i++)
        {
            if(a->next!=nullptr) 
                a=a->next;
            else 
                return nullptr;
        }
        while(a->next!=nullptr)
        {
            a=a->next;
            b=b->next;
        }
        return b;
    }
};
```