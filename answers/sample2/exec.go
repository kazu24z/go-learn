package sample2

import (
	"fmt"

	"go-learn/registry"
)

func Exec() {
	fmt.Println("executed from sample2.go")
}

func init() {
	registry.RegisterExec("sample2", Exec)
}
