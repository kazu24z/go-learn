package registry

import "fmt"

// 関数を動的に呼び出すためのマップ
var ExecMap = make(map[string]func())

// 関数を登録するための関数
func RegisterExec(name string, fn func()) {
    ExecMap[name] = fn
}

// 指定された名前の関数を実行する
func Exec(name string) error {
    fn, exists := ExecMap[name]
    if !exists {
        return fmt.Errorf("No such Exec function for %s", name)
    }
    fn()
    return nil
}
