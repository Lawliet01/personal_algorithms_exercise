package stack

import (
	"fmt"
	"testing"
)

func TestIsLinkedStackWork(t *testing.T) {
	stack := InitLinkedList()
	stack.push(123)
	stack.push("abc")
	stack.push("giao")
	stack.push(8999)
	stack.push(false)
	stack.pop()
	fmt.Println(stack)
}
