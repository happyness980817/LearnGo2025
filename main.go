package main

import (
	"fmt"
	"time"
)

// channel
// channel: goroutine 과 main 함수 사이에 (또는 goroutine 과 goroutine 사이 등등...) 정보를 전달하기 위한 방법

func main() {
	c := make(chan bool)
	people := [2]string{"john","jane"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println(person)
	c <- true
}