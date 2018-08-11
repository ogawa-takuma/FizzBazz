package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

func SrtStdin() (stringInput string) {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	return
}

func main() {

	fmt.Printf("ループ回数を整数で入力してください：")
	roopNum, err := strconv.Atoi(SrtStdin())

	if err != nil{
		fmt.Println("数値を入力してください")
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
