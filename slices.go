package main

import "fmt"

func main() {
	languages := [9]string{
		"c", "Lisp", "C++", "Java", "Python",
		"JavaScript", "Ruby", "Go", "Rust",
	}
	// fmt.Println(languages)

	// classics := languages[0:3]
	// fmt.Println(classics)

	// modern := make([]string, 4)
	modern := languages[3:7]
	fmt.Println("Modern languages:", modern)
}
