package main

import (
	"fmt"
)

func main() {
	var t, a, b int
	fmt.Scanf("%d\n", &t)
	c := 1
	for {
		fmt.Scanf("%d %d\n", &a, &b)	
		fmt.Printf("Case #%d: %d + %d = %d\n", c, a, b, a + b)
		t = t - 1
		c = c + 1
		if t == 0 {
			return
		}
	}
}