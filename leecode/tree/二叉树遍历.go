package main

import "fmt"

// 二叉树的数据结构
type TreeNode3 struct {
	Data  int
	Left  *TreeNode3
	Right *TreeNode3
}

// 二叉树的实现
type Tree struct {
	root *TreeNode3
}

// 添加数据
func (self *Tree) Add(data int) {
	var queue []*TreeNode3
	newNode := &TreeNode3{Data: data}
	if self.root == nil {
		self.root = newNode
	} else {
		queue = append(queue, self.root)
		for len(queue) != 0 {
			cur := queue[0]
			queue = append(queue[:0], queue[0+1:]...)
			// 往右树添加
			if data > cur.Data {
				if cur.Right == nil {
					fmt.Println(cur.Data)
					cur.Right = newNode
				} else {
					queue = append(queue, cur.Right)
				}
				// 往左数添加
			} else {
				if cur.Left == nil {
					fmt.Println(cur.Data)
					cur.Left = newNode
				} else {
					queue = append(queue, cur.Left)
				}
			}
		}
	}
}

// 添加数据
func (self *Tree) Add2(data int) {
	var queue []*TreeNode3
	newNode := &TreeNode3{Data: data}
	if self.root == nil {
		self.root = newNode
	} else {
		queue = append(queue, self.root)
		for len(queue) != 0 {
			cur := queue[0]
			queue = append(queue[:0], queue[0+1:]...)
			// 往右树添加
			if data > cur.Data {
				if cur.Right == nil {
					fmt.Println(cur.Data)
					cur.Right = newNode
				}
				// 往左数添加
			} else {
				if cur.Left == nil {
					fmt.Println(cur.Data)
					cur.Left = newNode
				} else {
					queue = append(queue, cur.Left)
				}
			}
		}
	}
}

// 前序遍历 根 ---> 左 --->右
func (self *Tree) preorderTraverse(node *TreeNode3) {
	if node == nil {
		return
	} else {
		fmt.Print(node.Data, " ")
		self.preorderTraverse(node.Left)
		self.preorderTraverse(node.Right)
	}

}

// 中序遍历 左 ---> 根 --->右
func (self *Tree) inorderTraverse(node *TreeNode3) {
	if node == nil {
		return
	} else {
		self.inorderTraverse(node.Left)
		fmt.Print(node.Data, " ")
		self.inorderTraverse(node.Right)

	}
}

// 后序遍历 左 ----> 右 ---> 根
func (self *Tree) postTraverse(node *TreeNode3) {
	if node == nil {
		return
	} else {
		self.postTraverse(node.Left)
		self.postTraverse(node.Right)
		fmt.Print(node.Data, " ")
	}
}

func main() {
	tree := &Tree{}
	tree.Add2(50)
	tree.Add2(45)
	tree.Add2(40)
	tree.Add2(48)
	tree.Add2(51)
	tree.Add2(61)
	tree.Add2(71)

	fmt.Println("前序遍历")
	tree.preorderTraverse(tree.root)
	fmt.Println("")
	fmt.Println("中序遍历")
	tree.inorderTraverse(tree.root)
	fmt.Println("")
	fmt.Println("后续遍历")
	tree.postTraverse(tree.root)
}
