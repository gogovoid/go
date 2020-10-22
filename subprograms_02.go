package main

import "fmt"

func sum(numbers ...int) int{
	r := 0
	for _ , number := range numbers {
		r = r + number
	}
	return r
}

func main(){
	//var inhaArray [7] string //array
//	var inhaSlice [] string //slice
//	inhaSlice = make([]string, 7)
	inhaSlice := []string{"i","n","h","a" }

//	inhaSlice[0] = "i"
//	inhaSlice[1] = "n"
//	inhaSlice[2] = "h"
//	inhaSlice[3] = "a"

	for i := 0 ; i < len(inhaSlice); i++ {
		fmt.Println(inhaSlice[i])
	}
	fmt.Println(sum(1,5,3))
	fmt.Println(sum(3,1))

//	fmt.Println(1)
}
