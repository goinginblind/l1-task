package main

import (
	"fmt"
)

// Разработать программу, которая в runtime способна определить тип переменной,
// переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).

func main() {
	all := getSomething()

	for _, v := range all {
		switch concValue := v.(type) {
		case int:
			mut := concValue + 1
			fmt.Printf("type: %T\n\twas: %v\n\tis:  %v\n\n", concValue, concValue, mut)
		case string:
			mut := "wow " + concValue + "!"
			fmt.Printf("type: %T\n\twas: %v\n\tis:  %v\n\n", concValue, concValue, mut)
		case bool:
			mut := !concValue
			fmt.Printf("type: %T\n\twas: %v\n\tis:  %v\n\n", concValue, concValue, mut)
		case chan any:
			fmt.Printf("type: %T\n\tchannel says: %v\n\n", concValue, <-concValue)
		}
	}
}

func getSomething() []any {
	var (
		myInt    = 123
		myString = "this is a string"
		myBool   = true
		myCh     = make(chan any, 1)
	)

	all := []any{myString, myInt, myBool, myCh}

	defer close(myCh)
	myCh <- "\"hi from channel!\""

	return all
}
