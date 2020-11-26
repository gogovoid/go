package main

import "fmt"

type subscriber struct {
	name   string
	rate   int
	active bool
}
func printInfo(s *subscriber) {
	fmt.Println("Name : ", s.name)
	fmt.Println("Monthly rate : ", s.rate)
	fmt.Println("Active? ", s.active)
}
func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5000
	s.active = true
	return &s
}
func applyDiscount(s *subscriber) {
	s.rate = 4000
}
func main() {
	s1 := defaultSubscriber("kim")
	//s1.rate = 4500
	applyDiscount(s1)
	printInfo(s1)

	s2 := defaultSubscriber("park")
	printInfo(s2)

	var s3 subscriber
	applyDiscount(&s3)
	fmt.Println(s3.rate)

}
