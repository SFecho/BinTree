package Stack

import "Util"


type ListNode struct{
	data Util.Object
	next * ListNode
}

type Stack struct{
	head * ListNode
	size int
	comp Util.Comparator
}

func New() * Stack{
	list := new(Stack)
	list.head = nil
	list.size = 0
	return list
}

func (this * Stack)Push(data Util.Object){
	node := &ListNode{data, this.head}
	this.size++;
	this.head = node
}

func (this * Stack) GetTop() Util.Object{
	if this.head != nil{
		return this.head.data
	}
	return nil
}

func (this * Stack)Empty() bool{
	return this.size == 0
}

func (this * Stack) Pop(){
	if this.head != nil {
		this.head = this.head.next
		this.size--
	}
}

func (this * Stack) GetSize() int{
	return this.size
}