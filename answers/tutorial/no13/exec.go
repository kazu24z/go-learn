package no13

import (
	"fmt"

	"go-learn/registry"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // keyがstring, valueがVertexのマップ。この時点はmにはnilが入っている（keyもない）

// Map
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no13/exec.go")
	//
	m = make(map[string]Vertex) // makeは初期化済みのマップを作成する
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func init() {
	registry.RegisterExec("tutorial/no13", Exec)
}
