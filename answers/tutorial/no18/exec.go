package no18

import (
	"fmt"

	"go-learn/registry"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Function_closures
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no18/exec.go")

	pos, neg := adder(), adder() // pos, negでそれぞれsumの状態を独立して持てる
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func init() {
	registry.RegisterExec("tutorial/no18", Exec)
}
