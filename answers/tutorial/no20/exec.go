package no20

import (
	"fmt"
	"math"

	"go-learn/registry"
)

type Vertex struct {
	X, Y float64
}

// func (レシーバー) 関数名() 戻り値の型型
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.Y + v.X*v.Y) // 平方根
}

// typeとmethod
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no20/exec.go")
	//
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

func init() {
	registry.RegisterExec("tutorial/no20", Exec)
}
