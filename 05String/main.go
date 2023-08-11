package main

import "fmt"

var num1 int
var num2 float32
var num3 float64

func main() {
	var str1 string = "Hello World"
	fmt.Println(str1)

	fmt.Printf("Length of str1: %v\n", len(str1))
	// mutating string
	str1 = "Hello World i am from victoria secret"
	fmt.Printf("Length of str1: %v\n", len(str1))

	num1 := 1000
	fmt.Println(num1) //shorthand declaration  it is an int , int holds 8 bytes in 64-bit processor
	num1 = 2000       //mutating num1
	fmt.Println(num1)
	num1 = 38388
	fmt.Println(num1)

}