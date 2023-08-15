// Q9 ->find the biggest and smallest value in the array}

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	arr := []int{3, 5, 7, 8, 2, 9, 4, 6, 13, -2, -1} //array declear and initialize

	BiggestN, SmallestN := FindBigestAndSmallest(arr)

	fmt.Println("The Biggest and Smallest Number of this array is : ", BiggestN, SmallestN)
	var arr2 [12]int
	Arrayval := randArray(arr2)
	BiggestN, SmallestN = FindBigestAndSmallest(Arrayval)
	fmt.Println("The Biggest and Smallest Number of this array is : ", BiggestN, SmallestN)

}

func randArray(arr2 [12]int) []int {

	for index, _ := range arr2 {

		arr2[index] = rand.Intn(30)

	}

	return arr2[:]

}
func FindBigestAndSmallest(arr []int) (int, int) {
	res1 := 0
	res2 := 100
	for _, val := range arr {
		if res1 < val {
			res1 = val
		}

		if res2 > val {
			res2 = val
		}

	}
	return res1, res2

}
