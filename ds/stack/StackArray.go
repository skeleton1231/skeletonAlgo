package stack

import "fmt"

type ArrayStack struct {
	//数据 infterface 代表数据类型
	data []interface{}
	//栈顶指针
	top int
}

// 创建ArrayStack
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// 使用多态实现StackInferface
func (as *ArrayStack) IsEmpty() bool {
	return as.top < 0
}

func (as *ArrayStack) Push(v interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top = 1
	}

	if as.top > len(as.data)-1 {
		as.data = append(as.data, v)
	} else {
		as.data[as.top] = v
	}
}

func (as *ArrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}

	v := as.data[as.top]
	as.top -= 1
	return v
}

func (as *ArrayStack) Top() interface{} {
	if as.IsEmpty() {
		return nil
	}

	return as.data[as.top]
}

func (as *ArrayStack) Flush() {
	as.top = -1
}

func (as *ArrayStack) Print() {
	if as.IsEmpty() {
		fmt.Println("The Stock is empty")
	} else {
		for i := as.top; i >= 0; i-- {
			fmt.Println(as.data[i])
		}
	}
}
