package main

import "fmt"

func main() {
	var RandomCountry = [...]string{
		"Guatemala",
		"Ethiopia",
		"Romania",
		"Australia",
	}

	for index, country := range RandomCountry {
		fmt.Println(index, country)
	}

}
