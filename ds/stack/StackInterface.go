package stack

type StackInterface interface {
	Push(v interface{})
	Pop()
	IsEmpty()
	Top() interface{}
	Flush()
}
