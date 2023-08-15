/* Q11-> Create and assign values to 2d array 3X3 array
store rows to columns and columns to rows in the array
Write a func to perform those operations
Input:

1 2 3

4 5 6

7 8 9
Output:

 1 4 7

 2 5 8

 3 6 9*/

package main

import "fmt"

func main() {
	size := 3
	arrVal := [3][3]int{{3, 7, 9}, {9, 5, 2}, {8, 7, 1}}

	ans := transposeOfMatrix(arrVal, size)

	count := 0
	fmt.Println("The transverse of matrix :->")
	for _, val := range ans {
		if count == 3 {
			fmt.Println("")
			count = 0
		} else {

			fmt.Println(val, " ")
			count++
		}
	}

}

func transposeOfMatrix(arr2 [3][3]int, size int) [3][3]int {

	var arr3 [3][3]int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			arr3[i][j] = arr2[j][i]
		}
	}
	return arr3

}
