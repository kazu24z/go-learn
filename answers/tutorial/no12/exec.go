package no12

import (
	"fmt"

	"go-learn/registry"
)

// Range_sliceやmapの要素を繰り返し処理する
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no12/exec.go")
	//
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v) // vはpow要素のコピーの値
	}
	fmt.Println("powの値は変わっていない", pow)

	// もし元の値に対して操作したい場合
	for index, value := range pow {
		pow[index] = pow[index] * 2
		fmt.Printf("%d*2 = %d\n", value, pow[index])
	}

	fmt.Println("元のpowが2倍になっている", pow)

	// indexやvalueの省略
	// , value は省略可能
	for i := range pow {
		fmt.Printf("index: %d\n", i)
	}

	// valueだけ使いたい
	for _, value := range pow {
		fmt.Printf("value: %d\n", value)
	}
}

func init() {
	registry.RegisterExec("tutorial/no12", Exec)
}
