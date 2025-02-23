package no5

import (
	"fmt"

	"go-learn/registry"
)

// 配列Slices
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no5/exec.go")

	//
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("元配列: ", primes)

	var s []int = primes[1:4] // primesの1番目から4番目までの要素を取得。この右辺の[1:4]はスライス式という
	fmt.Println("スライス: ", s)

	// スライスは元配列の参照のようなもの。
	// スライスの要素を変更すると元配列の要素も変更される。
	s[0] = 100 // スライスの0番目の要素を100に変更 → これは元配列の1番目の要素も100に変更される
	fmt.Println("元配列: ", primes)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func init() {
	registry.RegisterExec("tutorial/no5", Exec)
}
