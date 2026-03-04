/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiam := 0                                           // global pointer to max diameter
    dfs(root, &maxDiam)                                    // recursively calculate diameter 
    return maxDiam
}

func dfs(root *TreeNode, maxDiam *int) int {
    if root == nil { return 0 }
    left := dfs(root.Left, maxDiam)                         // left subtree length
    right := dfs(root.Right, maxDiam)                       // right subtree height
    if left + right > *maxDiam { *maxDiam = left + right }  // update global pointer of max diameter
    return max(left, right) + 1                             // return subtree height
}