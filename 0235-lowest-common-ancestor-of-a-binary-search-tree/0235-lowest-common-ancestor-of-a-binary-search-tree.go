/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {            // nil check
        return nil
    }
    if max(p.Val, q.Val) < root.Val {                   // if nodes in one subtree, recurse that subtree
        return lowestCommonAncestor(root.Left, p, q)
    } else if min(p.Val, q.Val) > root.Val {
        return lowestCommonAncestor(root.Right, p, q)
    }
    return root                                         // base case: split (LCA) found
}