package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	var iface1, iface2 any // it will be similar to NULL interface

	iface1 = 4467
	iface2 = 118

	if sum, err := add(iface1, iface2); err != nil {
		fmt.Println("We can't do, ", err)
	} else {
		fmt.Println("The sum is : ", sum)
	}

}

func add(a, b any) (float64, error) {
	sum := 0.0

	if a == nil || b == nil {
		return 0, fmt.Errorf("a or b is nil")
	}

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return 0, errors.New("a or b must be of same type")
	}

	switch a.(type) {
	case int:
		sum = float64(a.(int) + b.(int))
	case int8:
		sum = float64(a.(int8) + b.(int8))
	case int16:
		sum = float64(a.(int16) + b.(int16))
	case int32:
		sum = float64(a.(int32) + b.(int32))
	case int64:
		sum = float64(a.(int64) + b.(int64))
	case uint:
		sum = float64(a.(uint) + b.(uint))
	case uint8:
		sum = float64(a.(uint8) + b.(uint8))
	case uint16:
		sum = float64(a.(uint16) + b.(uint16))
	case uint32:
		sum = float64(a.(uint32) + b.(uint32))
	case uint64:
		sum = float64(a.(uint64) + b.(uint64))
	}

	return sum, nil
}
