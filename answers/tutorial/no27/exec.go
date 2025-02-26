package no27

import (
	"fmt"
	"strconv"
	"strings"

	"go-learn/registry"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	// {127, 0, 0, 1}
	var strArr [4]string // これは配列

	for i, num := range ip {
		strArr[i] = strconv.Itoa(int(num)) // numはbyte型 > int型にした上で、Int to ASCII
	}

	return strings.Join(strArr[:], ".") // strArrは配列なので、strArr[:]でスライスにしてJoinに渡す
}

// interfaceの例
func Exec() {
	fmt.Println("Go runs on ")
	defer fmt.Println("executed from tutorial/no27/exec.go")
	//
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func init() {
	registry.RegisterExec("tutorial/no27", Exec)
}
