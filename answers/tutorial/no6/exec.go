package no6

import (
	"fmt"

	"go-learn/registry"
)

// 配列Slices
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no6/exec.go")

	q := []int{2, 3, 5, 7, 11, 13} // {x, y, z, ...} の表現でスライスリテラル
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true} // 配列を作成し、それを参照するスライスのリテラル
	fmt.Println(r)

	r = append(r, false) // appendでスライスに要素を追加できる
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func init() {
	registry.RegisterExec("tutorial/no6", Exec)
}
