package main

import "fmt"

func main() {
	var RandomCountry = [4]string{
		"Guatemala",
		"Ethiopia",
		"Romania",
		"Australia",
	}

	for i := 0; i < len(RandomCountry); i++ {
		option := RandomCountry[i]
		fmt.Println(i, option)
	}
}
