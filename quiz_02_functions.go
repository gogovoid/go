package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func input() int {
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	in = strings.TrimSpace(in)
	number, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
func absolute(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
func power(b int, e int) int{
	r := 1
	for i := 1; i<=e ; i++{
		r = r * b
	}
	return r
}
func main() {
	for x := 1; true; x++ {
		fmt.Print("1) 절대값 2) 팩토리얼 3) 피보나치 4) 거듭제곱 5) 종료 : ")
		n := input()
		if n == 1 {
			fmt.Print("정수 입력(절대값 계산) : ")
			fmt.Println(absolute(input()))
		} else if n == 2 {
			fmt.Print("정수 입력(팩토리얼 계산) : ")
			fmt.Println(factorial(input()))
		} else if n == 3 {
			fmt.Print("정수 입력(피보나치 출력) : ")
			f := input()
			fmt.Println(fibonacci(f))
		} else if n == 4 {
			fmt.Print("Base 입력 : ")
			b := input()
			fmt.Print("Exponent 입력 : ")
			e := input()
			fmt.Println(power(b, e))
		} else if n == 5 {
			os.Exit(3)
		} else {
			fmt.Print("잘못 된 입력 값입니다. 1~5사이의 수를 입력하세요!")
		}

	}

}
