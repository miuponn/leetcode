/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    if root == nil { return 0 }                     // nil check for base case

    var dfs func(curr *TreeNode, maxVal int) int        
        dfs = func(curr *TreeNode, maxVal int) int {
        if curr == nil { return 0 }                 // nil check to exit early
        count := 0                                  // var for result (num of good nodes)
        if curr.Val >= maxVal {                     // process curr node = good node
            count = 1                               // increment count
            maxVal = curr.Val                       // update max for subtree
        }
        count += dfs(curr.Left, maxVal)             // recurse left sub, pass subtree max and return count
        count += dfs(curr.Right, maxVal)            // recurse right sub, pass subtree max and return count
        return count                                // base case: end of sub reached, return sub's count
    }
    return dfs(root, root.Val)                      // make call to dfs
}