package no10

import (
	"fmt"
	"strings"

	"go-learn/registry"
)

// 配列Slices_連想配列
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no10/exec.go")

	//
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[0][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][1] = "X"
	board[1][1] = "O"
	board[2][2] = "O"
	board[2][0] = "X"
	board[2][1] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func init() {
	registry.RegisterExec("tutorial/no10", Exec)
}
