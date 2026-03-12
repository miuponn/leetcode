/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    res := []int{root.Val}                                          // wrap result int in slice to mutate 

    var dfs func(node *TreeNode) int
        dfs = func(node *TreeNode) int {
            if node == nil { return 0 }                             // base case: end of subtree, add 0 to path sum

            // post-order: left + right subtree processing
            leftMax := dfs(node.Left)                               // process sum of left subtree
            leftMax = max(leftMax, 0)                               // negative single downward paths computed as 0

            rightMax := dfs(node.Right)                             // process sum of right subtree
            rightMax = max(rightMax, 0)                             // negative single downward paths computed as 0

            // possible path 1: node connects l + r subs
            res[0] = max(res[0], node.Val + leftMax + rightMax)     // only keep possible path 1 if beat max path so far

            // possible path 2: propagate only 1 sub up to parent extended path
            return node.Val + max(leftMax, rightMax)                // node + 1 sub propagated upwards to parent subtree sum   
        }
    dfs(root)
    return res[0]
}
