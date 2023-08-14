/*
  Q5 -Create variables of all number types . 14 datatypes including rune and byte

declare and assign values of all variables

perform addition and multiplecation, assign the result to a another variable, print it.

*/

package main

import "fmt"

func main() {

	var (
		num   int     = 1233
		num1  int8    = 100
		num2  int16   = 1235
		num3  int32   = 787
		num4  int64   = 9187
		num5  uint8   = 123
		num6  uint16  = 1234
		num7  uint32  = 657
		num8  uint64  = 6778
		_     rune    = 'A' // rune is alias of int32
		_     byte    = 'b' // byte is alias of  uint8
		num11 float32 = 123.9988
		num12 float64 = 766.78767677
	)
	// going to perform addition operation

	var add = num + int(num1) + int(num2) + int(num3) + int(num4) + int(num5) + int(num6) + int(num7) + int(num8) // here we convert all different type of data type in similar data type and then i perform addition
	fmt.Println("The sum of the value of all Datatype are  :", add)

	//going  to perform multiplication of different datatype

	var multiply = float32(num) * float32(num1) * float32(num2) * float32(num11) * float32(num12)

	fmt.Println("the multiplication of all the value of datatype are :", multiply)

}
