package main

import (
	"fmt"
)

type treenode struct {
	value       int
	left, right *treenode
}

func (p *treenode) xx() {
	if p == nil {
		return
	}
	fmt.Println(p.value)
	p.left.xx()
	p.right.xx()
}

func (node *treenode) gd() {
	if node == nil {
		return
	}
	//队列
	cs := 0
	result := []int{}
	nodes := []*treenode{node} //节点指针切片
	for len(nodes) > 0 {       //直到没有指针，nil
		cs++
		size := len(nodes)
		count := 0

		for count < size {
			count++
			curnode := nodes[0]
			result = append(result, curnode.value)
			nodes = nodes[1:] //pop出第0个
			if curnode.left != nil {
				nodes = append(nodes, curnode.left)
			}
			if curnode.right != nil {
				nodes = append(nodes, curnode.right)
			}
		}
	}

	fmt.Println("层数是：", cs)
	for _, v := range result {
		fmt.Println(v)
	}

}

func main() {
	root := &treenode{value: 1}
	root.left = &treenode{value: 2}
	root.right = &treenode{value: 3}
	root.left.left = &treenode{value: 5}
	root.left.right = &treenode{value: 6}
	root.right.left = &treenode{value: 7}
	root.right.right = &treenode{value: 8}
	root.right.left.left = &treenode{value: 9}

	root.gd()
	fmt.Println("-------")
	root.xx()
}
