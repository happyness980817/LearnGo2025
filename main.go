package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	foods := [2]string{"kimchi","pizza"}
	for _, food := range foods {
		go isDelicious(food, channel)
	}
	fmt.Println(<-channel)
	fmt.Println(<-channel)
} 

func isDelicious(food string, channel chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(food)
	channel <- true
}