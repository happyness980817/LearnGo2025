package main

import (
	"fmt"
	mydict "learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greetings!"
	err := dictionary.Add(word,definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("word:", word, "definition:", hello)
	err2 := dictionary.Add(word,definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}