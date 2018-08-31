package LinkedList

import "util"


type ListNode struct{
	data util.Object
	next * ListNode
}

type LinkedList struct{
	head * ListNode
	size int
	comp util.Comparator
}

func New(cmp util.Comparator) * LinkedList{
	list := new(LinkedList)
	list.head = nil
	list.size = 0
	list.comp = cmp
	return list
}

func (this * LinkedList)PushFront(data util.Object){
	node := &ListNode{data, this.head}
	this.size++;
	this.head = node
}

func (this  *  LinkedList)GetHead()* ListNode{
	return this.head
}

func (this * LinkedList)Release(data util.Object) bool{
	var del_ptr ** ListNode = &this.head

	for *del_ptr != nil{
		if this.comp((*del_ptr).data, data) == 0{
			*del_ptr = (*del_ptr).next
			this.size--
			return true;
		}
		del_ptr = &(*del_ptr).next
	}
	return false
}

func (this * LinkedList)Empty() bool{
	return this.size == 0
}

func (this * LinkedList)PopFront(){
	if this.head != nil {
		this.head = this.head.next
		this.size--
	}

}

func (this * LinkedList)Sort(){
	tmpHead := this.head
	this.head = nil
	var tail ** ListNode = nil
	for tmpHead != nil{
		curNode := tmpHead
		tmpHead = tmpHead.next
		tail = &this.head
		for *tail != nil && this.comp((*tail).data, curNode.data) == -1{
			tail = &(*tail).next
		}
		curNode.next = *tail
		*tail = curNode
	}
}

/*-------ListNode-------*/
func (this  *  ListNode)GetNext()* ListNode{
	return this.next
}

func (this  *  ListNode)GetData()util.Object{
	return this.data
}

func (this * LinkedList)Reverse(){
	tmpHead := this.head
	this.head = nil
	for tmpHead != nil{
		curNode := tmpHead
		tmpHead = tmpHead.next
		curNode.next = this.head
		this.head = curNode
	}
}
