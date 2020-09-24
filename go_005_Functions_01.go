package main

import(
	"fmt"
)

func calcArea(w float64, h float64){
	//var area float64
	area := w * h
	fmt.Printf("%5.2f\n",area/10.0)
}

func main(){
	calcArea(4.2, 3.0)
	calcArea(5.2, 3.5)
	calcArea(2.2, 2.2)
	/*
	var w, h, area float64
	w = 4.2
	h = 3.0
	area = w * h
	//fmt.Println(area/10.0)
	fmt.Printf("%5.2f\n",area/10.0)
	w = 5.2
	h = 3.5
	area = w * h

	fmt.Printf("%5.3f\n",area/10.0)
	fmt.Printf("%5d %s %t\n", 10, "go", true)
	fmt.Printf("%v %#v %T %%\n", 1.2, 1.2, 1.2)
	fmt.Printf("%v %#v %T %%\n", "\t", "\t", "\t")
	fmt.Printf("%v %#v %T %%\n", true, true, true)
	*/
}
