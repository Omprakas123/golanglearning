/*
	Q6. create variables of 14 types including rune, byte and also take 2 interface{} type variables;For one assign int and the other one assign float32 values.

assign values to them and perform addition and multiplication, print the details
*/
package main

import (
	"fmt"
	"reflect"
)

type DataType interface{} // declaration of empty interface it is alice of any
// type any interface{} // it's a empty interface in empty interface you can assign and data type of value
func main() {
	var num1 int = 10
	var num2 int8 = 127
	var num3 int16 = 5667
	var num4 int32 = -72877878
	var num5 int64 = 8767871389183189
	var num6 float32 = 866341.83919230138
	var num7 float64 = -5454545.3198

	var num10 byte = 243       //alias for uint8
	var n1um1 rune = 838713727 //alias for int32

	var num12 float32 = 324.1234
	var num13 float32 = 255
	var num14 float32 = 7
	var num15 float32 = 87237
	var num16 float32 = 2187

	//below are the interfaces
	var add_val DataType = num1 + int(num2) + int(num3) + int(num4) + int(num5) + int(num6) + int(num7) // here any is alias of empty interface

	var mul_val DataType = float32(num10) * float32(n1um1) * float32(num12) * float32(num13) * float32(num14) * float32(num15) * float32(num16)

	fmt.Println("The type of Variable add_val is : ", reflect.TypeOf(add_val), "  The type of Variable mulval is : ", reflect.TypeOf(mul_val))

	fmt.Println("The val all  variable after adding :", add_val, "The value of all varible after multiple :", mul_val)
}
