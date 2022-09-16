package linkedlist

import (
	"fmt"
)

// 节点
type Node struct {
	next  *Node
	value interface{}
}

/*
链表 head 代表第一个节点，不带头结点的链表，初始时 head 等于 null ; 带头结点的链表，初始时 head->next 等于 null ;
*/
type List struct {
	head *Node
	len  uint
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

func NewList() *List {
	return &List{NewNode(0), 0}
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) GetValue() interface{} {
	return n.value
}

// 在某个节点后面插入节点
func (list *List) InsertAferNode(n *Node, v interface{}) bool {

	if nil == n {
		return false
	}
	//
	newNode := NewNode(v)
	newNode.next, n.next = n.next, newNode
	list.len++
	return true
}

// 在某个节点前面插入
func (list *List) InsertBeforeNode(n *Node, v interface{}) bool {

	if nil == n && n == list.head {
		return false
	}

	current := list.head.next
	previous := list.head

	for nil != current {
		if current == n {
			break
		}
		previous = current
		current = current.next
	}

	if nil == current {
		return false
	}

	newNode := NewNode(v)
	previous.next = newNode
	newNode.next = current
	list.len++
	return true
}

func (list *List) FindNodeByIndex(index uint) *Node {
	if index > list.len {
		return nil
	}

	current := list.head.next
	var i uint = 0
	for ; i <= index; i++ {
		current = current.next
	}
	return current

}

// 在链表头部后面插入节点
func (list *List) InsertToHead(v interface{}) bool {
	return list.InsertAferNode(list.head, v)
}

// 插入到链表尾部后面插入节点
func (list *List) InsertToTail(v interface{}) bool {
	current := list.head
	for nil != current.next {
		current = current.next
	}
	return list.InsertAferNode(current, v)
}

// 遍历打印链表 实现Value Interface
func (list *List) PrintList() {
	current := list.head.next
	format := ""
	for nil != current {
		format += fmt.Sprintf("%v", current.GetValue())
		if nil != current {
			format += "->"
		}
	}
	fmt.Println(format)
}
