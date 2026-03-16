package main

import (
	"fmt"
	"time"
)

// channel
// channel: goroutine 과 main 함수 사이에 (또는 goroutine 과 goroutine 사이 등등...) 정보를 전달하기 위한 방법

func main() {
	c := make(chan string)
	people := [5]string{"john","jane","dal","japanguy","larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- person + " is sexy"
}