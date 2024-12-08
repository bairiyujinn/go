package main

import "fmt"

const max = 100

// BiNode结构体表示二叉树结点
type BiNode struct {
	Data   byte
	Lchild *BiNode
	Rchild *BiNode
}

// BiTree 是指向二叉树节点的指针类型，代表二叉树
type Bitree *BiNode

// Linkqueue 结构体模拟队列，用于层序遍历 二叉树时候存储节点
type Linkqueue struct {
	front int
	rear  int
	Data  []BiNode
}

// 构建二叉树，通过先序遍历的输入方式（@表示空节点）
func CreatBitree() *BiNode {
	var ch byte
	_, err := fmt.Scanf("%c", &ch)
	if err != nil {
		return nil
	}
	if ch == '@' {
		return nil
	} else {
		T := BiNode{Data: ch}
		T.Lchild = CreatBitree()
		T.Rchild = CreatBitree()
		return &T
	}
}

// 入队
func pushqueue(T BiNode, q *Linkqueue) {
	if (q.rear+1)%len(q.Data) != q.front {
		q.Data[q.rear] = T
		q.rear = (q.rear + 1) % len(q.Data)
	} else {
		fmt.Println("queue overflow")
	}
}

// 出队
func popqueue(q *Linkqueue) {
	if q.front != q.rear {
		q.front = (q.front + 1) % len(q.Data)
	} else {
		fmt.Println("queue underflow")
	}
}

// 判读是否队空
func emptyqueue(q Linkqueue) bool {
	return q.front == q.rear
}

// 层序遍历
func leave(T Bitree) {
	if T == nil {
		return
	}
	q := Linkqueue{
		front: 0,
		rear:  0,
		Data:  make([]BiNode, 100),
	}
	pushqueue(*T, &q)
	for !emptyqueue(q) {
		temp := &q.Data[q.front]
		fmt.Printf("%c", temp.Data)
		popqueue(&q)
		if temp.Lchild != nil {
			pushqueue(*temp.Lchild, &q)
		}
		if temp.Rchild != nil {
			pushqueue(*temp.Rchild, &q)
		}
	}
}

func main() {
	T := CreatBitree()
	leave(T)
}

/*
输出：ABCDEF
*/
