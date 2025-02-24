package no21

import (
	"fmt"
	"math"

	"go-learn/registry"
)

type Vertex struct {
	X, Y float64
}

// 値レシーバ = vは生成したインスタンスのコピーなので、値の変更不可
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.Y + v.X*v.Y) // 平方根
}

// ポインタレシーバ = 生成した構造体（インスタンス？）自身にアクセスできる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// typeとmethod
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no21/exec.go")
	//
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

func init() {
	registry.RegisterExec("tutorial/no21", Exec)
}
