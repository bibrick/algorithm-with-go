package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)

	var t int
	var s string
	fmt.Fscanf(reader, "%d\n", &t)


	for {
		fmt.Fscanf(reader, "%s\n", &s)
		check(s)

		t = t -1
		if t == 0 {
			return
		}
	}
}

func check(s string) {
	var stack []string
	for _, char := range s {
		if string(char) == "(" {
			stack = append(stack, string(char))
		} else if string(char) == ")" && len(stack) != 0{
			if stack[len(stack) - 1] == "(" {
				stack = stack[0:len(stack) - 1]
			}
		} else if string(char) == ")" && len(stack) == 0{
			fmt.Println("NO")
			return
		}
	}

	if len(stack) != 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}