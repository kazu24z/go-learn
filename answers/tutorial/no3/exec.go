package no3

import (
	"fmt"

	"go-learn/registry"
)

type Vertex struct {
	X, Y int
}

func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no3/exec.go")

	v1 := Vertex{1, 2} // X, Yの値を指定
	v2 := Vertex{X: 1} // Yはゼロ値が暗黙的に入る。 Name: 構文：フィールドの一部だけを指定
	v3 := Vertex{}     // X, Yはゼロ値が暗黙的に入る
	p := &Vertex{1, 2} // X=1, Y=2を値に持つVertexのポインタ（アドレス値）が入る

	fmt.Println(v1, v2, v3, p)
}

func init() {
	registry.RegisterExec("tutorial/no3", Exec)
}
