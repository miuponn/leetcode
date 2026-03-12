/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    parts := []string{}                                // slice of values to form string

    // pre-order: node, left, right
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            parts = append(parts, "null")              // if node null, add "null" to parts and exit
            return                           
        }
        parts = append(parts, strconv.Itoa(node.Val))  // else, convert ndoe value to str and add to parts
        dfs(node.Left)                                 // process left and right subtrees
        dfs(node.Right)
    }
    dfs(root)
    return strings.Join(parts, ",")                    // return slice as string with "," separator
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    vals := strings.Split(data, ",")                   // split data into slice of items
    i := 0                                             // global index to iterate vals

    // pre-order: node, left, right
    var dfs func() *TreeNode
    dfs = func() *TreeNode {
        if vals[i] == "null" {                         // "null" = nil node
            i++
            return nil
        }
        val, _ := strconv.Atoi(vals[i])                // convert item to int value, iterate next item
        i++
        node := &TreeNode{Val:val}                     // init node with value as Val
        node.Left = dfs()                              // process l and r subtree nodes next
        node.Right = dfs()
        return node                                    // return node after building its subtrees
    }
    return dfs()
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */