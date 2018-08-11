package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

/* FizzBuzzの要素を構造体にまとめる */
type roopElement struct {
	RoopNum int `json:"RoopNum"`
	Str3 string `json:"Str3"`
	Str5 string `json:"Str5"`
}

func main() {

	bytes, err := ioutil.ReadFile("roopElement.json")

	if err != nil {
		fmt.Println(err)
	}

	roopE := new(roopElement)
	if err := json.Unmarshal(bytes, &roopE); err != nil {
		fmt.Println(err)
	}

	for i := 1; i <= roopE.RoopNum; i++ {
		if i%15 == 0 {
			fmt.Println(roopE.Str3, roopE.Str5)
		} else if i%3 == 0 {
			fmt.Println(roopE.Str3)
		} else if i%5 == 0 {
			fmt.Println(roopE.Str5)
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
