package main

import (
	"fmt"
	mydict "learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first":"First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}