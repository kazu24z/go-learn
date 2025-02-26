package no22

import (
	"fmt"
	"math"

	"go-learn/registry"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	// tがnilの場合の制御を書いておく
	if t == nil {
		fmt.Println("nilです")
		return
	}
	fmt.Println("M():", t.S)
}

type F float64

func (f F) M() {
	fmt.Println("M():", f)
}

// Interface values
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no22/exec.go")
	//
	var i I // この時点ではnil

	var t *T // nil

	i = t // 値nil, 型*tのiを生成
	describe(i)
	i.M() // i自体はnilじゃないからM()を呼べるが、M()でnil処理してないとpanic

	// i は　(&T{"Hello"}, T)
	i = &T{"Hello"} // 実質的に、(値, 型)のタプル = 型情報とその型の値を持つ特殊なコンテナみたいなもの
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func init() {
	registry.RegisterExec("tutorial/no22", Exec)
}
