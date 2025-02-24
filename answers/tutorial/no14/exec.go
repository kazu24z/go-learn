package no14

import (
	"fmt"

	"go-learn/registry"
)

type Vertex struct {
	Lat, Long float64
}

// makeを使わずにmapリテラルで定義（初期化）
var m = map[string]Vertex{
	// "Bell Labs": Vertex{40.68433, -74.39967} のように型名{}だが、省略できる
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

// Map
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no14/exec.go")
	//
	fmt.Println(m)

	data, exist := m["Bell Labs"] // Goのmapにアクセスすると、keyが存在するかのフラグが返される
	fmt.Println("data:", data, "exist:", exist)
}

func init() {
	registry.RegisterExec("tutorial/no14", Exec)
}
