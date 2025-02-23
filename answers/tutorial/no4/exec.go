package no4

import (
	"fmt"

	"go-learn/registry"
)

// 配列Array
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no4/exec.go")

	var a [2]string // 要素数2のstring型配列
	a[0] = "Hello"
	a[1] = "World"
	// a[2] = "!" // コメントアウトを外すとエラー
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // 要素数6のint型配列
	fmt.Println(primes)
}

func init() {
	registry.RegisterExec("tutorial/no4", Exec)
}
