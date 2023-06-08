package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 21
	len := strconv.Itoa(x)
	s := len[:]
	fmt.Printf("len = %v, %T\n", s , s)
	
}