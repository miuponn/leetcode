/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    if root == nil { return 0 }

    var dfs func(curr *TreeNode, maxVal int) int
        dfs = func(curr *TreeNode, maxVal int) int {
        if curr == nil { return 0 }
        count := 0
        if curr.Val >= maxVal {
            count = 1
            maxVal = curr.Val
        }
        count += dfs(curr.Left, maxVal)
        count += dfs(curr.Right, maxVal)
        return count
    }
    return dfs(root, root.Val)
}