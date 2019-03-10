## 题目描述
输入一个递增排序的数组和一个数字S，在数组中查找两个数，使得他们的和正好是S，如果有多对数字的和等于S，输出两个数的乘积最小的。
## 输出描述:
对应每个测试案例，输出两个数，小的先输出。

## 思路
因为是递增排序的数组，所以先找个比s大的位置为r ，开始位置为l， 如果 a[ l ] + a[ r ]的 值 大于 S .那么r--, 小于 ，l++ ,等于输出
```c++
class Solution {
public:
    int aa(vector<int> &array,int &sum)
    {
        int n=array.size();
        int ans=-1;
        int l=0,r=n-1,mid;
        while(l<=r)
        {
            mid=(l+r)>>1;
            if(array[mid]>sum)
                r=mid-1;
            else 
            {
                ans=mid;
                l=mid+1;
            }
        }
        return ans;
    }
    vector<int> FindNumbersWithSum(vector<int> array,int sum) {
        vector<int> v;
        int l=0,r=aa(array,sum);
        while(l<r)
        {
            if(array[l]+array[r]==sum)
            {
                v.push_back(array[l]);
                v.push_back(array[r]);
                break;
            }
            else if(array[l]+array[r]>sum)
                r--;
            else 
                l++;
        }
        return v;
    }
};
```