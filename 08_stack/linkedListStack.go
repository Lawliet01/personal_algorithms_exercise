package stack

import (
	"fmt"
	"strings"
)

// 这里用了双向链表，其实不需要

type ListNode struct {
	value interface{}
	next  *ListNode
	prev  *ListNode
}

type LinkedListStack struct {
	tail   *ListNode
	length uint
}

func InitLinkedList() *LinkedListStack {
	return &LinkedListStack{
		tail:   nil,
		length: 0,
	}
}

func NewListNode(input interface{}) *ListNode {
	return &ListNode{
		value: input,
		next:  nil,
		prev:  nil,
	}
}

func (n *ListNode) getValue() interface{} {
	return n.value
}

func (list *LinkedListStack) push(input interface{}) {
	newNode := NewListNode(input)
	if list.length == 0 {
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.length++
	return
}

func (list *LinkedListStack) pop() {
	if list.length > 1 {
		list.tail.prev.next = nil
		list.tail = list.tail.prev
		list.length--
	} else if list.length == 1 {
		list.tail = nil
		list.length--
	}
}

func (list *LinkedListStack) String() string {
	var valueGroup []string
	cur := list.tail
	for cur != nil {
		valueGroup = append(valueGroup, fmt.Sprintf("%v", cur.getValue()))
		cur = cur.prev
	}
	return strings.Join(valueGroup, "->")
}
