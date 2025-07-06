package main

import "fmt"

func main() {
	a := 2
	b := &a // b 는 a 의 pointer 
	a = 10
	fmt.Println(a,b) // 10 0xc0000940a8
	fmt.Println(a,*b) // 10 10

	c := 1
	d := &c
	*d = 20
	fmt.Println(c) // 20
}
