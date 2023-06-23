package main

import "fmt"

func main() {
	mp := make(map[string]string)
	mp["123"] = "345"
	delete(mp, "123")
	fmt.Println(mp)
}