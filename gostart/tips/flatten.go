package main

import "fmt"

type Tree struct {
	value       int
	left, right *Tree
}
/*
将一个二叉树转换为单列右子树，类似链表
*/
func flatten(root Tree) {
	if root.left == nil {
		return
	}
	if root.left != nil {
		flatten(*root.left)
	}
	if root.right != nil {
		flatten(*root.right)
	}
	left := root.left
	right := root.right
	//左子树作为右子树，左子树抛弃
	root.left = nil
	root.right = left
	//将原先的右子树接到当前右子树的末端
	p := root
	if p.right != nil {
		p = *p.right
	}
	p.right = right
}

/*
遍历书并打印各元素
 */
func reversePrint(root Tree) {
	if &root == nil || &root.value == nil {
		return
	}
	fmt.Print(root.value)
	fmt.Print("-")
	reversePrint(*root.left)
	reversePrint(*root.right)
}

func main() {

	rr := Tree{
		value: 2,
		left:  nil,
		right: nil,
	}
	ll := Tree{
		value: 3,
		left:  nil,
		right: nil,
	}
	r := Tree{
		value: 2,
		right: &rr,
	}
	l := Tree{
		value: 3,
		left:  &ll,
	}

	root := Tree{
		value: 1,
		right: &r,
		left:  &l,
	}
	flatten(root)
	reversePrint(root)
}
