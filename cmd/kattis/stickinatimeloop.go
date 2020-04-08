package main

import (
	"fmt"
	"io"
)

func abs(x int64) int64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func main() {
	var a, counter int

	for i := 0; i < 1; i++ {
		_, err := fmt.Scanln(&a)
		if err == io.EOF {
			break
		}
	}

	counter = 1

	for i := 0; i < a; i++ {
		fmt.Printf("%d Abracadabra\n", counter)
		counter++
	}

}
