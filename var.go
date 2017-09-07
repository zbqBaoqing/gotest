package main

import (
	"errors"
	"fmt"
)

// go中全局变量的定义
var G1, G2, G3 string = "one", "tow", "three"

var (
	G4 int     = 12
	G5 bool    = false
	G6 float64 = 10.1102
)

// 常量定义

const (
	PI     = 3.1415926
	NUM    = 1000
	PREFIX = "HA-HA-"
)

//	:= 只能用在局部变量
//l1, l2, l3 := "ok", 110, true
func test() (string, error) {
	l1, l2, l3 := "ok", 110, true

	fmt.Printf("Global: G1:%#v  G2:%#v  G3:%#v  G4:%#v  G5:%#v  G6:%#v\n", G1, G2, G3, G4, G5, G6)

	fmt.Printf("Const: PI:%#v  NUM:%#v  PREFIX:%#v\n", PI, NUM, PREFIX)

	fmt.Printf("Local: l1:%#v  l2:%#v  l3:%#v\n", l1, l2, l3)

	G1 = "one --> ten"

	return G1, errors.New("no error. for test")
}

// 入口
func main() {
	if v, err := test(); err != nil {
		fmt.Printf("In main: return: %#v,  err:%s\n", v, err.Error())
	}
}
