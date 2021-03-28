package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var a, b int
		_, err := fmt.Scanf("%d %d\n", &a, &b)		
		if err == io.EOF{
			return
		}
		fmt.Printf("%d\n", a+b)
	}
}