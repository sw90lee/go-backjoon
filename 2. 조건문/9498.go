package main

import "fmt"

func main() {
	var a int

	fmt.Scanf("%d", &a)
	if a >= 90 && a <= 100 {
		fmt.Println("A")
	} else if a >= 80 && a <= 89 {
		fmt.Println("B")
	} else if a >= 70 && a <= 79 {
		fmt.Println("C")
	} else if a >= 60 && a <= 69 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
