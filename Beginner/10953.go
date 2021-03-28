package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	var a, b int
	for {
		fmt.Scanf("%d,%d\n", &a, &b)
		fmt.Printf("%d\n", a+b)
		t = t -1
		if t == 0 {
			return
		}
	}
}