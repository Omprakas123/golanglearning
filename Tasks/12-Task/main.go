/* Q12 -> Find the second biggest number in an array/slice

Find the second smallest number in an array/slice

Write everything using functions.

SecondBiggest(arr []int)int

SecondSmallest(arr []int)int */

package main

import "fmt"

func main() {

	slice1 := []int{2, 3, 12, 8, 9, 6, 4, 14} //short hand declearation of slice
	num1 := SecondBiggest(slice1)
	num2 := SecondSmallest(slice1)

	fmt.Println("The Secomd Greatest Number is:->", num1, " \nThe Secomd Loweast Number is :->", num2)
}

func SecondBiggest(slice []int) int {
	val1 := 0
	for _, val := range slice {

		if val > val1 {
			val1 = val
		}
	}
	sum2 := 0
	for _, val2 := range slice {

		if (sum2 < val2) && (val2 != val1) {
			sum2 = val2
		}
	}
	return sum2

}

func SecondSmallest(slice []int) int {
	val1 := 1000
	for _, val := range slice {

		if val < val1 {
			val1 = val
		}
	}
	sum2 := 1000
	for _, val2 := range slice {

		if sum2 > val2 && val2 != val1 {
			sum2 = val2
		}
	}
	return sum2
}
