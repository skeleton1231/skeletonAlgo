package queue

import "fmt"

type Node struct {
	v    interface{}
	next *Node
}

type QueueList struct {
	head *Node
	tail *Node
	len  int
}

func NewQueueList() *QueueList {
	return &QueueList{nil, nil, 0}
}

func (ql *QueueList) EnQueue(v interface{}) {
	node := &Node{v, nil}
	//入队时，tail->next= new_node, tail = tail->next (更新队尾下标)；
	if nil == ql.tail {
		ql.tail = node
		ql.head = node
	} else {
		ql.tail.next = node
		ql.tail = ql.tail.next
		/*
			ql.tail.next = node
			ql.tail = node
		*/
	}
	ql.len++
}

func (ql *QueueList) DeQueue() interface{} {
	//出队时，head = head->next。
	if ql.head == nil {
		return nil
	}
	v := ql.head.v
	ql.head = ql.head.next
	ql.len--
	return v
}

func (ql *QueueList) String() string {
	if ql.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := ql.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.v)
	}
	result += "<-tail"
	return result
}
