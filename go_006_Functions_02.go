package main

import "fmt"
import "errors"

func absolute(n int) int{
	if n<0 {
		return n * -1
	}else{
		return n
	}
}

func passFail(score float64) string{
	if score >= 60 {
		return "합격"
	}else if score < 0{
		err := errors.New("음수는 입력할 수 없습니다")
		return err.Error()
	}else{
		return "불합격"
	}
}

func manyReturn() (float64, bool, int){
	return 3.141592, false, 55
}

func factorial(n uint) uint{
	if n == 0{
		return 1
	}
	return n * factorial(n - 1)
}

func cArea(w float64, h float64)(normal float64, wrong error){
	if w < 0 {
		return 0, fmt.Errorf("유효하지 않은 너비 값 %0.2f", w)
	}
	if h < 0 {
		return 0, fmt.Errorf("유효하지 않은 높이 값 %0.2f", h)
	}
	area := w * h
	return area/10.0, nil
}

func main(){
	//a, err := cArea(3.3, -9.7)
	a, err := cArea(3.3, 9.7)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("면적 : %0.2f\n", a)
	}
//	fmt.Println(err)
//	fmt.Printf("%0.2f\n", a)


	fmt.Println(factorial(5))
	fmt.Println(absolute(-99))
	fmt.Println(passFail(59.7))
	fmt.Println(passFail(77.7))
	fmt.Println(passFail(-91.88))
	fmt.Println(absolute(17))
	myFloat, myBool, myInt := manyReturn()
	fmt.Println(myBool, myInt, myFloat)
	//fmt.Println(manyReturn())
}
