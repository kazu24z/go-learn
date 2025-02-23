package no7

import (
	"fmt"

	"go-learn/registry"
)

// 配列Slices_スライスリテラルとスライス式
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no7/exec.go")

	s := []int{2, 3, 5, 7, 11, 13} // スライスリテラル

	s = s[1:4] // スライスの1番目から4番目未満の要素を取得
	fmt.Println(s)

	s = s[:2] // スライスの最初から2番目未満の要素を取得
	fmt.Println(s)

	s = s[1:] // スライスの1番目から最後までの要素を取得
	fmt.Println(s)

	s = s[:] // スライスの最初から最後までの要素を取得
	fmt.Println(s)
}

func init() {
	registry.RegisterExec("tutorial/no7", Exec)
}
