package no19

import (
	"fmt"

	"go-learn/registry"
)

func fibonacci() func() int { // 「intを返す関数」を返す関数fibonacci
	//
	sum := 0
	n1, n2 := 0, 1
	return func() int {
		sum = n1
		n1, n2 = n2, n1+n2
		return sum
	}
}

// Function_closures
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no19/exec.go")

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func init() {
	registry.RegisterExec("tutorial/no19", Exec)
}
