package main

import "fmt"
import "os"

var LA int8 = 0

func error(c int8){
	fmt.Printf("Systax error : %c", c)
	os.Exit(1)// 종료
}

func match(token int8){
	if(LA == token){
		fmt.Scanf("%c", &LA)
	} else {
		error(token)
	}
}
func S(){
	if(LA == int8('(')){
		match(int8('(')); S(); match(int8(')')); S();
	}
}
func K(){
	if(LA == int8(-1)){

	} else {
		S(); match('\n'); K();
	}
}
func main(){
	fmt.Scanf("%c", &LA)
	K()
}
