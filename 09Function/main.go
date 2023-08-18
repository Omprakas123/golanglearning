package main

import "fmt"

func main() {
	greet()
	greet()
	greetby("harshini") //function call in golang
	greetbym("hey", "dileep")
	r := areaofrect(11.3, 12.4)
	fmt.Println("area of rectangular: ->", r)
}
func greet() { // function prototype
	fmt.Println("hello world") // function definition
}
func greetby(name string) {
	fmt.Println(name)
}
func greetbym(name, msg string) {
	fmt.Println(name, msg)
}
func areaofrect(l, b float32) float32 { //here float32 is an unamed return type
	return l * b
}

// func rect(l, b float32) (area, perimeter float32) {
// 	area = l * b
// 	perimeter = 2 * (l + b)
// 	return area, perimeter
// }
