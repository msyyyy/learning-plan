**二叉查找树** 

1、若它的左子树不为空，则左子树上所有的节点值都小于它的根节点值。

2、若它的右子树不为空，则右子树上所有的节点值均大于它的根节点值。

3、它的左右子树也分别可以充当为二叉查找树。

**二分查找**的思想，可以很快着找到目的节点，查找所需的最大次数等同于二叉查找树的高度。

缺点： 可能退化成链结构



**AVL平衡树**

1. 具有二叉查找树的全部特性。
2. 每个节点的左子树和右子树的高度差至多等于1。

插入

1、左-左型：做右旋。

2、右-右型：做左旋转。

3、左-右型：先做左旋，后做右旋。

4、右-左型：先做右旋，再做左旋。