package no8

import (
	"fmt"

	"go-learn/registry"
)

// 配列Slices_lengthとcapacity
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no8/exec.go")

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 長さを0にする
	s = s[:0]
	printSlice(s)

	// 長さを4にする（これ、元の配列を削除したわけではないから、元の配列の0番目から4番目未満の要素を取得しただけ）
	s = s[:4]
	printSlice(s)

	// 最初の2つの値を削除する
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func init() {
	registry.RegisterExec("tutorial/no8", Exec)
}
