/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isRoot (input *TreeNode) bool {
    if input.Left != nil || input.Right != nil {
        return true
    }
    return false
}

func swapLeaf (input *TreeNode) {
    if isRoot(input) {
        fmt.Println(input.Val, input.Right, input.Left)
        input.Left, input.Right = input.Right, input.Left
        fmt.Println(input.Val, input.Right, input.Left)
    } 
    return 
}

func invertTree(root *TreeNode) *TreeNode {
    if root != nil {
        swapLeaf(root)
        invertTree(root.Left)
        invertTree(root.Right)
    }
    return root
}
