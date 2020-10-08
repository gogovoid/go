package main

import (
	"fmt"
)

//var myScore struct{
//	name string
//	id int
//	grade float32
//}

//type score struct{
//	name string
//	id int
//	grade float32
//}
func main() {
	var subj string
	var prof string
	subjects := map[string]string{
		"pl": "김인하",
		"oop2": "이인하",
		"statistics": "최인하",
	}

	for {
		fmt.Println("과목명 : ")
		fmt.Scanln(&subj)
		fmt.Println("담당교수 : ")
		fmt.Scanln(&prof)

		value, contains := subjects[subj] // key check : string, bool
		if contains {
			if value == prof {
				fmt.Println("수강신청 되었습니다.")
				break
			} else {
				fmt.Println("교수명을 확인하세요.")
			}
		}else {
			fmt.Println("해당 과목은 존재하지 않습니다.")
		}
	}



	//var kim score
	//kim.name = "inha"
	//kim.id = 100
	//kim.grade = 3.71
	//fmt.Printf("%.1f\n", kim.grade)

	//id := 55
	//var pid *int
	//pid = &id
	//pid := &id
	//fmt.Println(id, &id)
	//fmt.Println(*pid, pid, &pid)
	//fmt.Println(reflect.TypeOf(*pid), reflect.TypeOf(pid))



	//fmt.Printf("%#v\n", myScore)

	//ranks := map[string]int{"gold":1, "bronze":3, "silver":2}
	//fmt.Println(ranks["bronze"])
	//
	//subjects := map[string]string{
	//	"pl": "김인하",
	//	"oop2": "이인하",
	//	"statistics": "최인하",
	//}
	//fmt.Println(subjects["pl"])

	//var s string = "123"
	//s := "123"
	//s = 55
	//fmt.Println(s)

	//for i, _ := range "이인하" {
	//	// foreach, for ... in ..., for (... : ...)
	//	fmt.Println(i)
	//}
	//fmt.Println(len("이인하"))
	//
	//for idx, r := range "123" {
	//	fmt.Println(idx, r)
	//}
	//fmt.Println(len("123"))

	//fmt.Println('A', unsafe.Sizeof('A'))
	//var a int = 7
	//var b float64 = 9.5

	//var c, d float64 = 4.2, 3.0
	//fmt.Println((c*d)/10.0)
	//fmt.Println(float64(a) + b) // Go는 묵시적 캐스팅을 불허한다
	/*
	fmt.Println(5 * '2')
	//fmt.Println(5 + true)
	//fmt.Println(false - 2)
	fmt.Println(5 + '2')
	 */
}