package main

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length int
}

func Initial(array []int) *Array {
	data := make([]int, len(array), len(array))
	copy(data, array)
	return &Array{
		data:   data,
		length: len(data),
	}
}

func (a *Array) len() int {
	return a.length
}

func (a *Array) isOutOfIndex(index int) bool {
	if index > cap(a.data)-1 {
		return true
	}
	return false
}

func (a *Array) Insert(input, index int) error {
	if a.isOutOfIndex(index) {
		return errors.New("out of Index")
	}
	a.data = append(a.data, input)
	a.length++
	for i := a.length - 1; i > index; i-- {
		a.data[i-1], a.data[i] = a.data[i], a.data[i-1]
	}
	return nil
}

func (a *Array) delete(index int) error {
	if a.isOutOfIndex(index) {
		return errors.New("out of Index")
	}
	for i := index; i < a.length-1; i++ {
		a.data[i], a.data[i+1] = a.data[i+1], a.data[i]
	}
	a.length--
	a.data = a.data[:a.length]
	return nil
}

func (a *Array) append(input int) {
	a.data = append(a.data, input)
	a.length++
}

func main() {
	newArray := Initial([]int{1, 2, 3, 4, 5})
	// fmt.Println(newArray.len())
	// fmt.Println(newArray.isOutOfIndex(6))
	// fmt.Println(newArray.isOutOfIndex(5))
	// if err := newArray.Insert(6, 5); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	newArray.append(6)
	if err := newArray.delete(3); err != nil {
		fmt.Println(err)
	}
	fmt.Println(newArray)
}
