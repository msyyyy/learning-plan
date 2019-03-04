## 题目描述
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数（时间复杂度应为O（1）

## 思路

建立一个辅助栈，专门存最小值。

```c++
class Solution {
public:
    stack<int> s1,s2;
    void push(int value) {
        s1.push(value);
        if(s2.empty()) s2.push(value);
        else
        {
            if(s2.top()<value)
                s2.push(s2.top());
            else 
                s2.push(value);
        }
    }
    void pop() {
        s1.pop();
        s2.pop();
    }
    int top() {
        return s1.top();
    }
    int min() {
        return s2.top();
    }
};
```