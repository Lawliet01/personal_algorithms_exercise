package stack

import (
	"fmt"
	"strings"
)

type arrayStack struct {
	data   []interface{}
	length uint
}

func Init() *arrayStack {
	return &arrayStack{
		data:   make([]interface{}, 0, 0),
		length: 0,
	}
}

func (s *arrayStack) push(v interface{}) {
	s.data = append(s.data, v)
	s.length++
}

func (s *arrayStack) pop() (r interface{}) {
	r = s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--
	return
}

func (s *arrayStack) String() string {
	var printOut []string
	for _, v := range s.data {
		printOut = append(printOut, fmt.Sprint(v))
	}
	return strings.Join(printOut, "->")
}

// func main() {
// 	stack := Init()
// 	stack.push(1)
// 	stack.push("b")
// 	stack.push("giao")
// 	fmt.Printf("%+v", stack)
// 	stack.pop()
// 	fmt.Printf("%+v", stack)

// 	// fmt.Println(popOut)
// }
