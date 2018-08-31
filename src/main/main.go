package main

import (
	"fmt"
	"util/BinTree"
	"util"
)

func comp(o1 util.Object, o2 util.Object)int{
	var a int = o1.(int)
	var b int = o2.(int)
	return a - b;
}

func main() {
	tree := BinTree.New(comp)
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	tree.InorderTraversal(
		func (data util.Object){
			fmt.Printf("%d\t", data)
		})
}
