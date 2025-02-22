package main

import (
    "fmt"
    "go-learn/answers"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <filename>")
        return
    }

    target := os.Args[1] // 実行時引数からファイル名を取得

    execFunc, exists := answers.ExecMap[target]
    if !exists {
        fmt.Println("Error: No such Exec function for", target)
        return
    }

    execFunc() // 指定された関数を実行
}
