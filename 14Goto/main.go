package main

import "fmt"

func main() {

	i := 1
loop:
	if i%2 == 0 {
		goto printme                      // when cursor will come to line  no 10 then it will atomatic goto statement printime
		//fmt.Println(i, "is an even number")
	}
	fmt.Println("om")
back:
	i++
	if i <= 10 {
		goto loop
	}

printme:
	if i <= 10 {
		fmt.Println(i, "is an even number")
		goto back
	}
}