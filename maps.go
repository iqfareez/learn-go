package main

import "fmt"

func main() {
	tahunLahir := map[string]int{
		"Abah": 1968, "Ibu": 1973, "Abg": 1998, "Aku": 2000, "Arif": 2003, "Haziq": 2009,
	}

	for name, tahun := range tahunLahir {
		fmt.Printf("%s lahir pd tahun %d\n", name, tahun)
	}
}
