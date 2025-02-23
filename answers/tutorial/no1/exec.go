package no1

import (
	"fmt"

	"go-learn/registry"
)

func Exec() {
	fmt.Print("Go runs on ")
	defer fmt.Println("executed from tutorial/no1/exec.go")

	i := 20
	fmt.Println("数値:", i)

	pi := &i // 変数iのアドレスをpiに格納
	fmt.Println("ポインタアドレス:", pi)

	var np *int                // ポインタ変数を宣言しただけなので...
	fmt.Println("ポインタ変数:", np) // 出力はnil

	np = pi
	fmt.Println("ポインタ変数:", np) // 出力はポインタアドレス

	j := *np
	fmt.Println("iのアドレスpiが入ったnpをデリファレンスすると...:", j) // アドレスからそのアドレスに格納された値を参照する

	j = 50
	fmt.Println("jの値を50に変更したらiは:", i, "jは:", j) // jはアドレスnpが参照している値(20)を渡しているだけなので、これを変更したところでiへの影響はない

	*np = 100
	fmt.Println("npの値を100に変更したら...:", i) // 出力は100
}

func init() {
	registry.RegisterExec("tutorial/no1", Exec)
}
