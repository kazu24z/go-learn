package no15

import (
	"fmt"

	"go-learn/registry"
)

// Map_操作
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no15/exec.go")
	//
	m := make(map[string]int) // 空のmap...この時点のmにm["hello"]としたらゼロ値0が入った状態になる
	// ※ m["hello"]はエラーにならないが存在しないので、m, exist = m["hello"]のexistはfalse
	fmt.Println(m)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")                    // delete()で指定したmapの指定したkeyを削除する
	fmt.Println("The value:", m["Answer"]) // 削除された要素にアクセスしてもゼロ値が返る。

	v, ok := m["Answer"] // 存在チェックokはfalse
	fmt.Println("The value:", v, "Present?", ok)
}

func init() {
	registry.RegisterExec("tutorial/no15", Exec)
}
