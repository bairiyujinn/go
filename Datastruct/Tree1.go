package main

import "fmt"

// TreeNode定义二叉树节点结构体
type TreeNode struct {
	Data  byte
	Left  *TreeNode
	Right *TreeNode
}

// PreorderTraversal先序遍历二叉树
func PreorderTraversal(root *TreeNode) {
	if root != nil {
		fmt.Printf("%c", root.Data)
		PreorderTraversal(root.Left)
		PreorderTraversal(root.Right)
	}
}

// InorderTraversal中序遍历二叉树
func InorderTraversal(root *TreeNode) {
	if root != nil {
		InorderTraversal(root.Left)
		fmt.Printf("%c", root.Data)
		InorderTraversal(root.Right)
	}
}

// PostorderTraversal后序遍历二叉树
func PostorderTraversal(root *TreeNode) {
	if root != nil {
		PostorderTraversal(root.Left)
		PostorderTraversal(root.Right)
		fmt.Printf("%c", root.Data)
	}
}

// CreateBinaryTree根据输入创建二叉树，用'@'表示空节点
func CreateBinaryTree() *TreeNode {
	var ch byte
	_, err := fmt.Scanf("%c", &ch)
	if err != nil {
		return nil
	}
	if ch == '@' {
		return nil
	}
	root := &TreeNode{Data: ch}
	root.Left = CreateBinaryTree()
	root.Right = CreateBinaryTree()
	return root
}

// Depth计算二叉树的深度
func Depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := Depth(root.Left)
	rightDepth := Depth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// LeafCount计算二叉树的叶子节点数量
func LeafCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return LeafCount(root.Left) + LeafCount(root.Right)
}

func main() {
	root := CreateBinaryTree()
	fmt.Println("PreOrder:")
	PreorderTraversal(root)
	fmt.Println()
	fmt.Println("InOrder:")
	InorderTraversal(root)
	fmt.Println()
	fmt.Println("PostOrder:")
	PostorderTraversal(root)
	fmt.Println()
	height := Depth(root)
	num := LeafCount(root)
	fmt.Printf("Depth:%d\n", height)
	fmt.Printf("Leaf:%d\n", num)
}
