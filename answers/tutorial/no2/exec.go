package no2

import (
	"fmt"

	"go-learn/registry"
)

type Vertex struct {
	X int // フィールド
	Y int
}

func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no2/exec.go")

	v := Vertex{1, 2} // フィールドに値を代入
	fmt.Println(v.X, v.Y)

	pv := &v
	pv.X = 1e9 // 10の9乗。普通に考えたら、(*pv).Xとするが、手間なのでpv.XでもフィールドXにもアクセスできる
	v.X = 0    // さらに通常のvからXに代入もできる
	fmt.Println(v)
}

func init() {
	registry.RegisterExec("tutorial/no2", Exec)
}
