package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Please give numbers to get average from (min 2): ")

	var input string
	fmt.Scanln(&input)

	var sum float64
	var n int

	for {

		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)

		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is", sum/float64(n))
}
