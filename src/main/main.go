package main

import (
	"fmt"
	"Util/BinTree"
	"Util"
	"math/rand"
	"time"
)

func comp(o1 Util.Object, o2 Util.Object)int{
	var a int = o1.(int)
	var b int = o2.(int)
	return a - b;
}

func main() {
	tree := BinTree.New(comp)
	//tree.Insert(4)
	//tree.Insert(5)
	//tree.Insert(2)
	//tree.Insert(3)
	//tree.Insert(1)

	var arr[20] int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 20; i++{
		arr[i] = rand.Intn(100)
		tree.Insert(arr[i])
	}

	tree.InorderTraversal(
		func (data Util.Object){
			fmt.Printf("%d\t", data)
		})
	tree.Delete(arr[rand.Intn(20)])
	fmt.Printf("\n")
	tree.InorderTraversal(
		func (data Util.Object){
			fmt.Printf("%d\t", data)
		})

	tree.Delete(arr[rand.Intn(20)])
	fmt.Printf("\n")
	tree.InorderTraversal(
		func (data Util.Object){
			fmt.Printf("%d\t", data)
		})
	fmt.Printf("\n")
	tree.PreorderTraversal(
		func (data Util.Object){
			fmt.Printf("%d\t", data)
		})

	fmt.Printf("\n")

	fmt.Printf("height = %d\n", tree.GetHeight())

	tree.PostorderTraversal(func (data Util.Object){
		fmt.Printf("%d\t", data)
	})
}
