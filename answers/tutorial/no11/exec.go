package no11

import (
	"fmt"

	"go-learn/registry"
)

// 配列Slices_
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no11/exec.go")

	var s []int // nilスライス
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0) // スライスの末尾に要素を追加
	printSlice(s)    // len=1 cap=1 [0]
	fmt.Printf("Address after append: %p\n", s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s) // len=2 cap=2 [0 1]
	fmt.Printf("Address after append: %p\n", s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4) // 一度に複数の要素を追加できる
	printSlice(s)          // len=5 cap=5 [0 1 2 3 4]
	fmt.Printf("Address after append: %p\n", s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func init() {
	registry.RegisterExec("tutorial/no11", Exec)
}
