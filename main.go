package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	var roopNum int
	roopNum, _ = strconv.Atoi(os.Args[1])

	for i := 1; i <= roopNum; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
