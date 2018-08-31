package BinTree

import (
	"Util/Stack"
	"Util"
	"Util/Queue"
)


type TreeNode struct{
	data  Util.Object
	left  * TreeNode
	right * TreeNode
}

type BinTree struct{
	root * TreeNode
	size int
	comp Util.Comparator
}

func (this * BinTree)GetHeight() int{
	queue := Queue.New()

	var level_last_node_ptr * TreeNode = this.root
	//var cur_ptr * TreeNode = nil

	var cur_queue_tail_ptr  * TreeNode = nil

	queue.Push(this.root)

	var level int = 0

	for queue.Empty() == false{
		cur_ptr := queue.GetFront().(* TreeNode)
		queue.Pop()
		if cur_ptr.left != nil{
			cur_queue_tail_ptr = cur_ptr.left
			queue.Push(cur_queue_tail_ptr)
		}
		if cur_ptr.right != nil{
			cur_queue_tail_ptr = cur_ptr.right
			queue.Push(cur_queue_tail_ptr)
		}
		if level_last_node_ptr == cur_ptr{
			level++
			level_last_node_ptr = cur_queue_tail_ptr
		}
	}
	return level
}

func New(comp Util.Comparator) *BinTree{
	tree := &BinTree{nil, 0, comp}
	return tree
}

func(this * BinTree) Insert(data Util.Object)bool{
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

func(this * BinTree) InorderTraversal(display func(data Util.Object)){
	stk_ptr := Stack.New()
	search_ptr := this.root
	for search_ptr != nil || stk_ptr.Empty() == false {
		for search_ptr != nil{
			stk_ptr.Push(search_ptr)
			search_ptr = search_ptr.left
		}
		if stk_ptr.Empty() == false{
			search_ptr = stk_ptr.GetTop().(* TreeNode)
			display(search_ptr.data)
			stk_ptr.Pop()
			search_ptr = search_ptr.right
		}
	}
}

func(this * BinTree) PreorderTraversal(display func(data Util.Object)){
	stk_ptr := Stack.New()
	search_ptr := this.root
	for search_ptr != nil || stk_ptr.Empty() == false {
		for search_ptr != nil{
			display(search_ptr.data)
			stk_ptr.Push(search_ptr)
			search_ptr = search_ptr.left
		}
		if stk_ptr.Empty() == false{
			search_ptr = stk_ptr.GetTop().(* TreeNode)

			stk_ptr.Pop()
			search_ptr = search_ptr.right
		}
	}
}

func (this * BinTree) PostorderTraversal(function func(data Util.Object)) {
	stack_ptr := Stack.New()
	var prev * TreeNode = nil
	search_ptr := this.root

	for search_ptr != nil || stack_ptr.Empty() == false{
		if search_ptr != nil{
			stack_ptr.Push(search_ptr)
			search_ptr = search_ptr.left
		}else{
			search_ptr = stack_ptr.GetTop().(* TreeNode)
			if search_ptr.right != nil && search_ptr.right != prev{
				search_ptr = search_ptr.right
				stack_ptr.Push(search_ptr)
				search_ptr = search_ptr.left
			}else{
				search_ptr = stack_ptr.GetTop().(* TreeNode)
				stack_ptr.Pop()
				function(search_ptr.data)
				prev = search_ptr
				search_ptr = nil
			}
		}

	}
}

func(this * BinTree) Delete(data Util.Object){
	 del_node := &this.root
	 for *del_node != nil{
	 	if this.comp(data, (*del_node).data) < 0{
			 del_node = &(*del_node).left
		}else if this.comp(data, (*del_node).data) > 0{
			del_node = &(*del_node).right
		}else{
			//case 1： 若当前删除节点的左儿子没节点
			if (*del_node).left == nil {
				*del_node = (*del_node).right
			}else if (*del_node).right == nil{
				*del_node = (*del_node).left
			}else{
				//找比当前要删除的值稍大的节点
				tmp := &(*del_node).right
				for (*tmp).left != nil{
					tmp = &(*tmp).left
				}
				new_val := *tmp
				(*del_node).data = new_val.data
				*tmp = new_val.right
			}
		}
	 }
}

