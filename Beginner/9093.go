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
	// var stack []string
	for {
		// var s string
		// fmt.Fscanln(reader, &s)
		// input, _, _ := reader.ReadLine()
		input, _ := reader.ReadString('\n')
		// input = strings.TrimSpace(input)
		for i, s := range input {
			fmt.Fprintf(writer, "%d: %s\n", i, string(s))
		}
		writer.Flush()
		t = t - 1
		if t == 0 {
			return
		}
	}
}
