package main

import (
	"fmt"
	"os"
	"strings"

	_ "go-learn/answers"
	_ "go-learn/answers/tutorial"
	"go-learn/registry"
)

// パス入力のエイリアス
var Prefixes = map[string]string{
	"t-": "tutorial/",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	target := parseAlias(os.Args[1])

	execFunc, exists := registry.ExecMap[target]
	if !exists {
		fmt.Println("Error: No such Exec function for", target)
		return
	}

	execFunc()
}

// 実行エイリアスを解釈する関数
func parseAlias(arg string) string {
	for alias, replacement := range Prefixes {
		if strings.HasPrefix(arg, alias) {
			return replacement + strings.TrimPrefix(arg, alias)
		}
	}
	return arg
}
