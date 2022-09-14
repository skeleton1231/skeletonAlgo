package linkedlist

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

func (list *List) InsertBeforeNode(n *Node, v interface{}) bool {

	if nil == n && n == list.head {
		return false
	}

	return true
}
