package no26

import (
	"fmt"

	"go-learn/registry"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// interfaceの例
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no26/exec.go")
	//
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	// 解説
	/**
	* fmtパッケージには以下が定義されている
	*
	* type Stringer interface{
	* 	String() string
	* }
	*
	* no26側でfunc Sting() stringの実装をしており、fmt.Println()したら、実装したString()を呼び出す
	 */
}

func init() {
	registry.RegisterExec("tutorial/no26", Exec)
}
