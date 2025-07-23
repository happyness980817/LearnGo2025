package main

import (
	"fmt"
	"learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "FirstWord"
	dictionary.Add(baseWord,"First word")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	err2 := dictionary.Delete("NotExistingWord")
	fmt.Println(err2)
}

