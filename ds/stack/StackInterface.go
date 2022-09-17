package stack

type StackInterface interface {
	Push(v interface{})
	Pop() interface{}
	IsEmpty()
	Top() interface{}
	Flush()
}
