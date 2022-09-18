package queue

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
