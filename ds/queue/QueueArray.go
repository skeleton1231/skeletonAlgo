package queue

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

func (ar *ArrayQueue) EnQueue(v interface{}) bool {
	if ar.tail == ar.capacity {
		return false
	}

	ar.q[ar.tail] = v
	ar.tail++
	return true
}

func (ar *ArrayQueue) EnQueueV2(v interface{}) bool {
	if ar.tail == ar.capacity {
		if ar.head == 0 {
			return false
		}
		//数据搬移
		for i := ar.head; i < ar.tail; i++ {
			ar.q[i-ar.head] = ar.q[i]
		}
		// 搬移完之后重新更新head和tail
		ar.tail -= ar.head
		ar.head = 0
	}

	ar.q[ar.tail] = v
	ar.tail++
	return true

}

func (ar *ArrayQueue) DeQueue() interface{} {
	if ar.head == ar.capacity {
		return nil
	}

	v := ar.q[ar.head]
	ar.head++
	return v
}

func (ar *ArrayQueue) String() string {
	if ar.head == ar.tail {
		return "empty queue"
	}
	result := "head"
	for i := ar.head; i <= ar.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", ar.q[i])
	}
	result += "<-tail"
	return result
}
