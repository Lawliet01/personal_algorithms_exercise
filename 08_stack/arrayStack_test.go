package stack

import (
	"fmt"
	"testing"
)

func TestIsWork(t *testing.T) {
	stack := Init()
	stack.push(1)
	stack.push("b")
	stack.push("giao")
	fmt.Println(stack)
	// t.Logf("%+v", stack)
	stack.pop()
	fmt.Println(stack)

	// t.Logf("%+v", stack)

	// fmt.Println(popOut)
}
