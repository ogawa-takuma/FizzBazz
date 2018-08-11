package main

import "fmt"


func main() {
	var i int
	for i = 0; i < 100; i++ {
		if (i  == 0){
			fmt.Printf("%d\n", i)
		} else if (i % 3 == 0 && i % 5 == 0){
			fmt.Println("fizzbazz")
		} else if (i % 3 == 0){
			fmt.Println("fizz")
		} else if (i % 5 == 0){
			fmt.Println("bazz")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
