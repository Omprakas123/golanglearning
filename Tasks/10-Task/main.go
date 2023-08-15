/*Q10-> dimentional array 3 X 3 and assign values .

Get the sume of Each row and each column data

Write a func to perform those operations

Input:

1 2 3

4 5 6

7 8 9
Output:

Row 1 Sum: 6

Row 2 Sum: 15

Row 3 Sum: 24

Column 1 Sum: 12

Column 2 Sum: 15

Column 3 Sum: 18 */

package main

import "fmt"

func main() {
	size := 3
	var arr1 = [][]int{{2, 3, 4}, {5, 6, 7}, {7, 6, 5}, {5, 8, 9}} //initialize of

	SumRowAndColumn(arr1, size)

}
func SumRowAndColumn(arr2 [][]int, size int) {
	sum := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += arr2[i][j]
		}
		fmt.Println("The sum of row no", i+1, "is", sum)
		sum = 0
	}
	sum = 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sum += arr2[j][i]
		}
		fmt.Println("The sum of Column no", i+1, "is", sum)
		sum = 0
	}
}
