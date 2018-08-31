package Queue

import "Util"

type QueueNode struct {
	data Util.Object
	prev * QueueNode
	next * QueueNode
}

type Queue struct{
	head * QueueNode
	tail * QueueNode
	size int
}

func New() * Queue{
	queue := &Queue{nil,nil, 0}
	return queue
}

func(this * Queue) Push(data Util.Object){
	new_node := &QueueNode{data, nil, nil}
	if this.head == nil{
		this.head = new_node
		this.tail = new_node
	}else{
		this.tail.next = new_node
		new_node.prev = this.tail
		this.tail = new_node
	}
	this.size++
}

func(this * Queue) Pop(){
	if this.head != nil{
		this.head = this.head.next
		//如果出队的恰好为最后一个节点，则重置tail
		if this.head == nil {
			this.tail = nil
		}
		this.size--
	}
}

func (this * Queue)GetFront() Util.Object{
	if this.head != nil{
		return this.head.data
	}
	return nil
}

func (this * Queue)Empty()bool{
	return this.size == 0
}

func (this * Queue)GetSize() int{
	return this.size
}