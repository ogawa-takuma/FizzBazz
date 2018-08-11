package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strconv"
)


/* FizzBuzzの要素を構造体にまとめる */
type roopElement struct {
	RoopNum int `json:"RoopNum"`
	Str3 string `json:"Str3"`
	Str5 string `json:"Str5"`
}

var roopE = new(roopElement)

func FizzBuzz(roopNum int) string {
	
	if roopNum%15 == 0 {
		return roopE.Str3 + roopE.Str5
	} else if roopNum%3 == 0 {
		return roopE.Str3
	} else if roopNum%5 == 0 {
		return roopE.Str5
	} else {
		return strconv.Itoa(roopNum)
	}
}

func main() {

	bytes, err := ioutil.ReadFile("roopElement.json")

	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(bytes, &roopE); err != nil {
		fmt.Println(err)
	}

	for i := 1; i <= roopE.RoopNum; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
