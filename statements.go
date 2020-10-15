package main

import (
	"fmt"
)

func main() {
	pi := 3.141592
	radius := 10.0
	var a float32 = 10.3
	var b float64 = 11.7
	fmt.Println(float64(a) + b)
	area := int(radius * radius * pi)
	fmt.Println(area)

	if area >= 300 {
		fmt.Println("크다")
	} else {
		fmt.Println("작다")
	}

	var score int = 100;
	switch score/10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default :
		fmt.Println("F")
	}

	//for i:=0; i<3; i++ {
	//	fmt.Println("안녕")
	//}
	i := 0
	for i<3 {
		fmt.Println("안녕")
		i++
	}
}