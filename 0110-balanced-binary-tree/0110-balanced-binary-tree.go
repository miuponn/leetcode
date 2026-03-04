/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return dfs(root) != -1
}

func dfs(root *TreeNode) int {          // return -1 (unbalanced), 0 (no height added) or height
    if root == nil { return 0 }         // end of subtree, no height
    left := dfs(root.Left)              // traverse left subtree
    if left == -1 { return -1 }         // if unbalance in left sub, dont process abs and height
    right := dfs(root.Right)            // traverse right subtree
    if right == -1 { return -1 }        // if unbalance in right sub, dont process abs and height
    if abs(left - right) > 1 {          // calculate difference, return -1 if unbalanced
        return -1
    }
    return max(left, right) + 1         // return height of subtree
}

func abs(x int) int {                   // abs function helper
    if x < 0 { return -x }
    return x
}