package answers

import (
    "go-learn/answers/sample"
    "go-learn/answers/sample2"
)

// 関数を動的に呼び出すためのマップ
var ExecMap = map[string]func(){
    "sample":  sample.Exec,
    "sample2": sample2.Exec,
}

// sample.go の Exec
func ExecSample() {
    sample.Exec()
}

// sample2.go の Exec
func ExecSample2() {
    sample2.Exec()
}
