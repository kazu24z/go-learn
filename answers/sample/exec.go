package sample

import (
	"fmt"

	"go-learn/registry"
)

func Exec() {
	fmt.Println("executed from sample.go")
}

func init() {
	registry.RegisterExec("sample", Exec)
}
