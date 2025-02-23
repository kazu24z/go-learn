package main

import (
	"fmt"
	"os"

	_ "go-learn/answers"
	"go-learn/registry"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	target := os.Args[1]

	execFunc, exists := registry.ExecMap[target]
	if !exists {
		fmt.Println("Error: No such Exec function for", target)
		return
	}

	execFunc()
}
