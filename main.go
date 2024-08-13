package main

import (
	"fmt"
	"tdd/fn"
)

func main() {
	str := "//;\n1;2"
	sm, er := fn.Add(str)
	if er != nil {
		panic(er)
	}
	fmt.Println(sm)
}
