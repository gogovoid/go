package main
import "fmt"
import "reflect"
func main(){
	//var id int = 55
	//var pid *int
	//pid = &id
	id := 55
	grade := 3.91
	pid := &id
	pgrade := &grade
	fmt.Println(grade, &grade, reflect.TypeOf(grade))
	fmt.Println(*pgrade, pgrade, reflect.TypeOf(pgrade))
	fmt.Println(id, &id)
	fmt.Println(*pid, pid, &pid)
}
/*
package main
import "fmt"
func main(){
	number := 3
	cube(number)
	fmt.Println(number)
}
func cube(n int){
	n = n * n * n
	//fmt.Println(n)
}
*/
