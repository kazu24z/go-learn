package no1

import (
	"fmt"

	"go-learn/registry"
)

func Exec() {
	fmt.Println("executed from tutorial/no1/exec.go")
}

func init() {
	registry.RegisterExec("tutorial/no1", Exec)
}
