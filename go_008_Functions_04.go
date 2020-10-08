
package main
import "fmt"
import "errors"
func divide(dividend float64, divisor float64)(float64, error) {
	if divisor == 0.0{
		return 0, errors.New("0으로 나눌 수 없음")
	}
	return dividend / divisor, nil
}
func main(){
	var n1, n2 float64
	fmt.Scanf("%f %f", &n1, &n2)
	quotient, err := divide(n1, n2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}

/*
package main
import "fmt"
func testPointer() *float32 {
	var f =  float32(2.71)
	return &f
}
func main(){
	var pf *float32 = testPointer()
	fmt.Println(*pf)
}
*/

/*
package main
import "fmt"
func main(){
	number := 3
	cube(&number)
	fmt.Println(number)
}
func cube(n *int){
	*n = *n * *n * *n
	//fmt.Println(n)
}
*/
