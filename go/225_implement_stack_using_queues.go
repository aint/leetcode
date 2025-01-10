package main

import "fmt"

type MyStack struct {
	data []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *MyStack) Pop() int {
	l := len(s.data)
	if l == 0 {
		return 0
	}
	r := s.data[l-1]
	if l == 1 {
		s.data = nil
	} else {
		s.data = s.data[:l-1]
	}
	return r
}

func (s *MyStack) Top() int {
	l := len(s.data)
	if l == 0 {
		return 0
	}
	return s.data[l-1]
}

func (s *MyStack) Empty() bool {
	return s.data == nil
}

func testMyStack() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	fmt.Println("2 =", s.Pop())
	fmt.Println("1 =", s.Top())
	fmt.Println("false =", s.Empty())
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
