## 题目描述
输入n个整数，找出其中最小的K个数。例如输入4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4,。

## 思路

1. 通过类似快排的方法

2. 用set

```c++
// 类似快排 
class Solution {
public:
    void aa(vector<int> &in, int &k,int begin,int r)
    {
        int l=begin;
        while(l<r)
        {
            while(l<r&&in[r]>=in[begin])
                r--;
            while(l<r&&in[l]<=in[begin])
                l++;
            swap(in[l],in[r]);
        }
        swap(in[l],in[begin]);
        if(l==k-1||l==k) return;
        if(l>=k-1) 
            aa(in,k,begin,l);
        else if(l<k-1)
            aa(in,k,l+1,r);
        return;
    }
    vector<int> GetLeastNumbers_Solution(vector<int> input, int k) {
        vector<int> v;
        if(input.size()==0||input.size()<k) return v;
        else if(input.size()==1)
        {
            v.push_back(input[0]);
            return v;
        }
        aa(input,k,0,input.size()-1);
        for(int i=0;i<k;i++)
        {
            v.push_back(input[i]);
        }
        return v;
    }
};
```