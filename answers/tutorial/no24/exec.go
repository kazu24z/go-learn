package no24

import (
	"fmt"

	"go-learn/registry"
)

// 型アサーション
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no24/exec.go")
	//
	var i interface{} = "hello"

	s := i.(string) // 空インターフェースに入れた値は.(string)のような型アサーションが必要
	fmt.Println(s)

	s, ok := i.(string) // 型アサーションは指定した型を持っているかのboolを返す
	fmt.Println(s, ok)

	f, ok := i.(float64) // ちゃんとboolを受け取っていれば違っていてもゼロ値返すでpanicならない
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

func init() {
	registry.RegisterExec("tutorial/no24", Exec)
}
