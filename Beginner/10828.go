package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Scanf("%d\n", &t)
	var s []int
	for {
		var cmd string
		var num int
		fmt.Fscanf(reader, "%s %d\n", &cmd, &num)
		if cmd == "push" {
			s = append(s, num)
		} else if cmd == "pop" {
			if len(s) == 0 {
				fmt.Println(-1)
			} else {
				xs := s
				s = s[:len(s) - 1]
				fmt.Println(xs[len(xs) - 1]	)
			}
		} else if cmd == "size" {
			fmt.Println(len(s))
		} else if cmd == "empty"{
			if len(s) == 0 {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		} else if cmd == "top" {
			if len(s) == 0 {
				fmt.Println(-1)
			} else {
				fmt.Println(s[len(s) - 1])
			}
		}
		// fmt.Println(s)
		t = t - 1
		if t == 0 {
			return
		}
	}
}