/*
 ----> Q4. Create global variables and local variables with the same name in different scopes
           Identify which variable values are taken within or outside of scopes

*/

package main

import "fmt"

var temp1 int = 4

func temp() {

	fmt.Println("temp1: ", temp1) //global variables can be accessed from anywhere
}

func main() {
	temp()
	temp1 = 56 //global variables can be accessed from anywhere
	fmt.Println("temp1: ", temp1)
	var temp2 = 48
	fmt.Println("temp2: ", temp2)
}
