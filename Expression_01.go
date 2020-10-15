package main

import "fmt"

var a int = 1
func f() int{
	a = 2
	return 5;
}
func main() {
	fmt.Println(a + f())
}