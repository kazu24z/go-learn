package no23

import (
	"fmt"

	"go-learn/registry"
)

// Interface values
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no23/exec.go")
	//
	var i interface{} // メソッドを持たないインターフェースは空インターフェース
	describe(i)

	i = 42 // 空インターフェースは任意の型の値を保持できる
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func init() {
	registry.RegisterExec("tutorial/no23", Exec)
}
