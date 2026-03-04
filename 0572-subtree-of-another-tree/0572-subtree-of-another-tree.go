/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {                        
    if root == nil { return false }                                            // base case
    if subRoot == nil { return true }                                          // nil tree always a subtree
    if root.Val == subRoot.Val {                                               // if roots match, check if sametree
        if isSameTree(root, subRoot) { return true }
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)     // recurse into left and right subtrees
}

func isSameTree(root1 *TreeNode, root2 *TreeNode) bool {
    // if nodes are both null, return true
    if root1 == nil && root2 == nil { return true }
    // mismatch found, return false
    if (root1 == nil && root2 != nil) || (root2 == nil && root1 != nil ) || (root1.Val != root2.Val) {
        return false
    }
    return isSameTree(root1.Left, root2.Left) && isSameTree(root1.Right, root2.Right)
}