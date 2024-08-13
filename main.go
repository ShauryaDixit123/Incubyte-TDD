package main

import (
	"fmt"
	"tdd/fn"
)

func main() {
	str := "//;\n1;2" // simple check case added to check if program is working or not.
	sm, er := fn.Add(str)
	if er != nil {
		panic(er)
	}
	fmt.Println(sm)
}

// Add is the sum function which accepts numbered string saperated
// all testcases are present in test folder
// index_test.go

// COMMAND TO RUN TESTCASES:----->
// cd test
// go test -bench=.
