package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputstr string

	for i := 0; i < 1; i++ {
		consoleReader := bufio.NewReader(os.Stdin)
		input, _ := consoleReader.ReadString('\n')
		inputstr = input
	}

	arrstr := strings.Split(inputstr, " ")

	var thisstr = ""
	counterarr := len(arrstr)
	//var result = "no"
	//var countersame int
	//var thispos int
	//countersame = counterarr
	for i := 0; i < counterarr; i++ {
		thisstr = arrstr[i]
		//thispos = i

		//fmt.Println("Hel" == "Hel")

		fmt.Println(thisstr)
		for j := i; j < counterarr; j++ {
			if i != j {
				//fmt.Println("-- " + arrstr[j])
				//fmt.Println("-- " + thisstr + " x " + arrstr[j] + " -- ")
				//fmt.Println(strings.Compare(strings.TrimSpace(thisstr), strings.TrimSpace(arrstr[j])))
				if strings.TrimSpace(arrstr[j]) == "Hel" {
					fmt.Printf("Loop %d\n", j)
					fmt.Println(arrstr[j])
				}
			}
			//fmt.Printf("Pos %d\n", thispos)

		}
	}
}
