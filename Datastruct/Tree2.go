package main

import "fmt"

// TreeNode定义二叉树节点结构体
type TreeNode struct {
	Data  byte
	Left  *TreeNode
	Right *TreeNode
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

func midorder(root *TreeNode) {
	st := make([]*TreeNode, 0)
	for root != nil || len(st) > 0 {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}
		if len(st) > 0 {
			root = st[len(st)-1]
			st = st[:len(st)-1]
			fmt.Printf("%c", root.Data)
			root = root.Right
		}
	}
}

func main() {
	root := CreateBinaryTree()
	midorder((root))
	fmt.Println()
}

/*
中序遍历非递归输出二叉树
输入二叉树的扩展的先序遍历序列，建立一棵二叉树，输出中序遍历序列（要求非递归算法）。
二叉树结点类型为char,特殊字符为@。
输入先序遍历序列：ABD@F@@@CE@@@
输出二叉树的中序遍历序列为：DFBAEC
*/
