package main

import "fmt"

func calculateTotal(x []int) (result int) {

	if len(x) == 0 {
		result = 0
	} else {
		for _, v := range x {
			result += v
		}
	}
	return
}

func main() {
	numbers := []int{12, 34, 0, 8, 23}
	fmt.Println(calculateTotal(numbers))
}
