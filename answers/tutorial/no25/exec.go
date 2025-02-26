package no25

import (
	"fmt"

	"go-learn/registry"
)

func do(i interface{}) {
	switch v := i.(type) { // 型Switchのアサーションはtype
	case int: // caseは型を指定. 通ったcaseの型に自動でキャストされる
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// 型switch
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no25/exec.go")
	//
	do(21)
	do("hello")
	do(true)
}

func init() {
	registry.RegisterExec("tutorial/no25", Exec)
}
