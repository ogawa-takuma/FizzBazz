package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

/* FizzBuzzの要素を構造体にまとめる */
type roopElement struct {
	roopNum int
	str3 string
	str5 string
}

func SrtStdin() (stringInput string) {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	return
}

func main() {

	roopE := roopElement{str3:"Fizz", str5:"Buzz"}
	
	var err error

	fmt.Printf("ループ回数を整数で入力してください：")
	roopE.roopNum, err = strconv.Atoi(SrtStdin())

	if err != nil{
		fmt.Println("数値を入力してください")
		os.Exit(0)
	}

	for i := 1; i <= roopE.roopNum; i++ {
		if i%15 == 0 {
			fmt.Println(roopE.str3, roopE.str5)
		} else if i%3 == 0 {
			fmt.Println(roopE.str3)
		} else if i%5 == 0 {
			fmt.Println(roopE.str5)
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
