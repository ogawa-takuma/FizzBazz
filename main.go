package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	
	roopNum, err := strconv.Atoi(os.Args[1])

	if err != nil{
		fmt.Println("第一引数は数値を入力してください")
		os.Exit(0)
	}

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
