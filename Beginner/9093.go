package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var t int
	fmt.Fscanf(reader, "%d\n", &t)
	var stack []string
	for {
		input, _, _ := reader.ReadLine()
		input = append(input, '\n')
		for _, s := range input {
			if string(s) == " " || string(s) == "\n" {
				for j, _ := range stack {
					fmt.Fprint(writer, string(stack[len(stack)-(j+1)]))
				}
				fmt.Fprintf(writer, " ")
				writer.Flush()
				stack = nil
			} else {
				stack = append(stack, string(s))
			}
		}
		fmt.Fprintf(writer, "\n")
		writer.Flush()
		t = t - 1
		if t == 0 {
			return
		}
	}
}
