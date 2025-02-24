package no16

import (
	"encoding/json"
	"fmt"
	"strings"

	"go-learn/registry"
)

// Map_文字列に出現する単語の出現回数を数える
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no16/exec.go")

	for index, sentence := range list {
		jsonData, err := json.MarshalIndent(WordCount(sentence), "", " ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println(index+1, "回目")
		fmt.Println(string(jsonData))
	}
}

func WordCount(s string) map[string]int {
	// 末尾の.や!をとる
	extract := s[:len(s)-1] // len(s)だと、0からsの文字数分-1　つまり最初から最後までなので、そこから1引くことで末尾除去

	// extractをスペース区切りでスプリットしてスライスに格納
	words := strings.Split(extract, " ")

	// 以降ループ処理
	// 初回：スプリットをkey, 出現回数をvalueにするmapを作成
	// 2回目以降：既存のkeyのvalueを+1する
	m := map[string]int{}

	for _, word := range words {
		if _, exist := m[word]; !exist {
			m[word] = 0
		}
		m[word] += 1
	}

	// {"I": 1, "am": 1, "Tom": 1} のような出力
	return m
}

func init() {
	registry.RegisterExec("tutorial/no16", Exec)
}

// テストデータ
var list = []string{
	"I am Tom Tom Tom.",
	"He likes playing tennis.",
	"She want to be pianist",
}
