package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	var c float64 = float64(a) / float64(b)
	fmt.Println(c)
}
