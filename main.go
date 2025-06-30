package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int, uppercase string){
	defer fmt.Println("I'm done.") 
	// defer - funtion 끝나고 해당 코드 실행
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, up := lenAndUpper("nico")
	fmt.Println(totalLength,up)
}