package main

import "fmt"

/*
单链表基本操作
author:leo
*/

type ListNode struct {
	v    interface{}
	next *ListNode
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		v:    v,
		next: nil,
	}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.v
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   nil,
		length: 0,
	}
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	// fmt.Println(v)
	newNode := &ListNode{
		v:    v,
		next: nil,
	}
	// fmt.Printf("newNode1{v:%v,next:%v}\n", newNode.v, newNode.next)

	if p.next != nil {
		newNode.next = p.next
		// fmt.Printf("newNode{v:%v,next:%v}\n", newNode.v, newNode.next)
	}
	p.next = newNode
	this.length++
	return true
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	newNode := &ListNode{
		v:    v,
		next: nil,
	}
	if p == this.head {
		newNode.next = this.head
		this.head = newNode
		this.length++
		return true
	}

	cur := this.head
	for cur.next != nil {
		if cur.next == p {
			newNode.next = cur.next
			cur.next = newNode
			this.length++
			return true
		}
		cur = cur.next
	}
	return false
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	// fmt.Println(v)
	newNode := &ListNode{
		v:    v,
		next: nil,
	}

	if this.head != nil {
		newNode.next = this.head
	}
	this.head = newNode
	// fmt.Printf("newNode1{v:%v,next:%v}\n", newNode.v, newNode.next)

	this.length++
	return true
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	newNode := &ListNode{
		v:    v,
		next: nil,
	}
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = newNode
	this.length++
	return true
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index > this.length-1 {
		return nil
	}
	cur := this.head
	for i := uint(0); i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	cur := this.head

	if cur == p {
		this.head = this.head.next
		this.length--
		return true
	}

	for cur.next != nil {
		if cur.next == p {
			cur.next = cur.next.next
			this.length--
			return true
		}
	}
	return false
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}

func main() {
	list := NewLinkedList()
	// node1 := NewListNode(123)
	// node2 := NewListNode("abc")
	// node3 := NewListNode(true)
	// fmt.Println(node1, node2, node3)
	list.InsertToHead("start")
	list.InsertAfter(list.head, "abc")
	list.InsertAfter(list.head, 123)
	list.InsertAfter(list.head, false)
	list.InsertBefore(list.head, "newStart")
	list.InsertToTail("kkk")
	list.InsertToTail(23432)
	list.DeleteNode(list.head)
	list.DeleteNode(list.head)
	node := list.FindByIndex(3)
	// fmt.Println(list)
	list.Print()
	fmt.Printf("%+v", node)
}
