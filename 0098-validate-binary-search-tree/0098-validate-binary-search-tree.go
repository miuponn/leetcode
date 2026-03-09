/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil { return true }

    var dfs func(curr *TreeNode, min int, max int) bool
        dfs = func(curr *TreeNode, min int, max int) bool {
            if curr == nil { return true }
            if curr.Val <= min || curr.Val >= max { 
                return false 
            }
            return dfs(curr.Left, min, curr.Val) && dfs(curr.Right, curr.Val, max)
        }
    return dfs(root, math.MinInt, math.MaxInt)
}