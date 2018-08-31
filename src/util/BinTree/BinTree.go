package BinTree

import (
	"util/LinkedList"
	"util"
)


type TreeNode struct{
	data util.Object
	left * TreeNode
	right * TreeNode
}

type BinTree struct{
	root * TreeNode
	size int
	comp util.Comparator
}

func New(comp func(o1 util.Object, o2 util.Object) int) *BinTree{
	tree := &BinTree{nil, 0, comp}
	return tree
}

func(this * BinTree) Insert(data util.Object)bool{
	new_node := new(TreeNode)
	new_node.left = nil
	new_node.right = nil
	new_node.data = data

	ins_ptr := &this.root

	for *ins_ptr != nil{
		if this.comp(new_node.data, (*ins_ptr).data) < 0{
			ins_ptr = &(*ins_ptr).left
		}else if this.comp(new_node.data, (*ins_ptr).data) > 0 {
			ins_ptr = &(*ins_ptr).right
		}else{
			return false
		}
	}
	*ins_ptr = new_node
	this.size++
	return true
}

func(this * BinTree) InorderTraversal(display func(data util.Object)){
	stk_ptr := LinkedList.New(this.comp)
	search_ptr := this.root
	for search_ptr != nil || stk_ptr.Empty() == false {
		for search_ptr != nil{
			stk_ptr.PushFront(search_ptr)
			search_ptr = search_ptr.left
		}
		if stk_ptr.Empty() == false{
			search_ptr = stk_ptr.GetHead().GetData().(* TreeNode)
			display(search_ptr.data)
			stk_ptr.PopFront()
			search_ptr = search_ptr.right
		}
	}
}