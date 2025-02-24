package no17

import (
	"fmt"
	"math"

	"go-learn/registry"
)

// fn → 引数, func(float64, float64) float64　→ 引数が2つのfloat64, 戻りがfloat64のfunc型
// この関数は 関数を引数に取って、3, 4を渡して実行させる
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no17/exec.go")

	// hypotには右辺で定義した関数を入れることができる
	// 要は無名関数(jsの cost hypot = ()=> {return math.sqrt()} と同じ)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y) // 引数の平方根 169→13
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow)) // math.Pow関数そのものを渡している
}

func init() {
	registry.RegisterExec("tutorial/no17", Exec)
}
