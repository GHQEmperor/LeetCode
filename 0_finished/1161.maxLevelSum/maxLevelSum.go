package maxlevelsum

import "math"

/*
给你一个二叉树的根节点root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。

请你找出层内元素之和 最大 的那几层（可能只有一层）的层号，并返回其中最小 的那个。

示例 1：
输入：root = [1,7,0,7,-8,null,null]
输出：2
解释：
第 1 层各元素之和为 1，
第 2 层各元素之和为 7 + 0 = 7，
第 3 层各元素之和为 7 + -8 = -1，
所以我们返回第 2 层的层号，它的层内元素之和最大。

示例 2：
输入：root = [989,null,10250,98693,-89388,null,null,null,-32127]
输出：2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	c := &count{}
	c.DFS(root, 0)
	idx, max := 0, math.MinInt64
	for i := len(*c) - 1; i >= 0; i-- {
		if (*c)[i] >= max {
			idx, max = i, (*c)[i]
		}
	}
	return idx + 1
}

type count []int

func (c *count) DFS(node *TreeNode, deep int) {
	if node == nil {
		return
	}
	for len(*c) <= deep {
		*c = append(*c, 0)
	}
	(*c)[deep] += node.Val
	c.DFS(node.Left, deep+1)
	c.DFS(node.Right, deep+1)
}
