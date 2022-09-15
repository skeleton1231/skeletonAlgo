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
