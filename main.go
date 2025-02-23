package main

import (
	"fmt"
	"os"
	"strings"

	_ "go-learn/answers"
	_ "go-learn/answers/tutorial"
	"go-learn/registry"
)

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
	prefixes := map[string]string{
		"t-": "tutorial/",
	}

	for alias, replacement := range prefixes {
		if strings.HasPrefix(arg, alias) {
			return replacement + strings.TrimPrefix(arg, alias)
		}
	}
	return arg
}
