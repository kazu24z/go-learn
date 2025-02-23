package no9

import (
	"fmt"

	"go-learn/registry"
)

// 配列Slices_makeで作成
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no9/exec.go")

	//
	a := make([]int, 5) // 0値で初期化された長さ5のスライスを作成
	printSlice("a", a)

	b := make([]int, 0, 5) // 長さ0で容量5のスライスを作成
	printSlice("b", b)

	c := b[:2] // bの0~1番目の要素を取得
	printSlice("c", c)

	d := c[3:5] // TODO: slice[start:end]すると、cap(d)は「元のスライスの容量-start」になる=5-3=2
	printSlice("d", d)

	b = b[:cap(b)] // bの容量=5を取得し、0~4番目の要素を取得（ゼロ値で初期化されたスライス）
	printSlice("b", b)

	b = b[:1] // bの0番目の要素を取得
	printSlice("b", b)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func init() {
	registry.RegisterExec("tutorial/no9", Exec)
}
