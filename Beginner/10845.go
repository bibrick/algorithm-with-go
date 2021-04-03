package main

import(
	"fmt"
	"os"
	"bufio"
)


func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanf(reader, "%d\n", &t)
	var a string
	var b int
	var q []int

	for {
		fmt.Fscanf(reader, "%s %d\n", &a, &b)
		if a == "push" {
			q = append(q, b)
		} else if a == "front" {
			if len(q) == 0 {
				fmt.Fprintf(writer, "%d\n", -1)
			} else {
				fmt.Fprintf(writer, "%d\n", q[0])
			}
		} else if a == "back" {
			if len(q) == 0 {
				fmt.Fprintf(writer, "%d\n", -1)
			} else {
				fmt.Fprintf(writer, "%d\n", q[len(q)-1])
			}
		} else if a == "empty" {	
			if len(q) == 0 {
				fmt.Fprintf(writer, "%d\n", 1)
			} else {
				fmt.Fprintf(writer, "%d\n", 0)
			}
		} else if a == "pop" {
			if len(q) == 0 {
				fmt.Fprintf(writer, "%d\n", -1)
			} else {
				fmt.Fprintf(writer, "%d\n", q[0])
				q = q[1:]
			}
		} else if a == "size" {
			fmt.Fprintf(writer, "%d\n", len(q))
		}
		writer.Flush()
		t = t - 1
		if t == 0 {
			return
		}
	}
	fmt.Fprintf(writer, "hello\n")
	writer.Flush()
}