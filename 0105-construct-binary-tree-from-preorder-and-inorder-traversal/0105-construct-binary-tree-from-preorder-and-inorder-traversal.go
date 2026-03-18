/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    preIndex, inIndex := 0, 0                      // indices to iterate both arrays

    var dfs func(int) *TreeNode
    dfs = func(limit int) *TreeNode {
        if preIndex >= len(preorder) {             // base case 1: reached end of tree
            return nil  
        }
        if inorder[inIndex] == limit {             // base case 2: end of building sub, continue in next sub
            inIndex++
            return nil
        }
        root := &TreeNode{Val: preorder[preIndex]} // init current node in preorder array
        preIndex++                                 // continue building in left subtree until base case
        root.Left = dfs(root.Val)
        root.Right = dfs(limit)                    // continue building in right subtree 

        return root
    }
    return dfs(math.MaxInt)                        // use very large num for initial limit
}