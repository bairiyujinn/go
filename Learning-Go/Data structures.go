package main
import "fmt"

// 结点
type btNode struct {
	Data   interface{}
	lchild *btNode
	rchild *btNode
}
// 二叉树
type biTree struct {
	root *btNode
}


//二叉树的创建
func main() {
	var tree *biTree
	var node1, node2, node3, node4 *btNode
	tree = new(biTree)

	node1 = &btNode{Data: 1}
	node2 = &btNode{Data: 2}
	node3 = &btNode{Data: 3}
	node4 = &btNode{Data: 4}

	tree.root = node1
	node1.lchild = node2
	node1.rchild = node3
	node2.lchild = node4

	fmt.Println(tree.root.Data)
	fmt.Println(tree.root.lchild.Data)
	fmt.Println(tree.root.rchild.Data)
	fmt.Println(tree.root.lchild.lchild.Data)

}
/*
输出：
1
2
3
4
*/



//层次遍历
func (bt *biTree) levelorder() []interface{} {
	var list []interface{}
	cur := bt.root
	if cur == nil {
		return list
	}
	Queue := []*btNode{cur}
	for len(Queue) > 0 {
		cur := Queue[0]
		Queue = Queue[1:]
		list = append(list, cur.Data)
		if cur.lchild != nil {
			Queue = append(Queue, cur.lchild)
		}
		if cur.rchild != nil {
			Queue = append(Queue, cur.rchild)
		}
	}
	return list
}

//创建二叉树
func creatTree(list []interface{}) *biTree {
	btree := &biTree{}
	if len(list) > 0 {
		btree.root = &btNode{Data: list[0]} //根节点
		for _, data := range list[1:] {
			btree.appendNode(data)  
		}
	}
	return btree
}

//添加子节点
func (bt *biTree) appendNode(data interface{}) {
	cur := bt.root
	if cur == nil {
		bt = &biTree{}
		return
	}
	Queue := []*btNode{cur}
	for len(Queue) > 0 {
		cur := Queue[0]
		Queue = Queue[1:]
		if cur.lchild != nil {
			Queue = append(Queue, cur.lchild)
		} else {
			cur.lchild = &btNode{Data: data}
			return
		}
		if cur.rchild != nil {
			Queue = append(Queue, cur.rchild)
		} else {
			cur.rchild = &btNode{Data: data}
			break
		}
	}
}

func main() {
	list := []interface{}{1, 2, 3, 4, 5, 6}
	tree := creatTree(list)
	fmt.Println(tree.levelorder())

	tree.appendNode(0)
	fmt.Println(tree.levelorder())
}
/*
输出：
[1 2 3 4 5 6]
[1 2 3 4 5 6 0]
*/
