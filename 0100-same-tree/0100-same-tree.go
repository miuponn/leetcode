/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    // if nodes both null, return true
    if p == nil && q == nil {                                                
        return true 
    }
    // mismatch found, return false
    if (p == nil && q != nil) || (q == nil && p != nil ) || (p.Val != q.Val) {
        return false
    }
    // recursively compare nodes in left subtree then right subtree, return true if both match
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}