package main

import "fmt"

func main() {
	a := []string{"Bob", "Michael", "Peter Gregory", "Al", "Edgar", "Habtu Gebrahed Ab", "Steve"}
	greatest := ""
	for i := 0; i < len(a)-1; i++ {
		if len(a[i]) >= len(a[i+1]) {
			greatest = a[i]
		} else {
			greatest = a[i+1]
		}
	}
	fmt.Println(greatest, "is the largest name in the list!")
	for i := 0; i < len(a); i++ {
		if a[i] == "Peter Gregory" {
			fmt.Println("Peter Gregory was included too. That's funny!")
		}
	}

}
