package main

import (
	"fmt"
	"time"
)

func main() {
	go deliciousCount("kimchi")
	deliciousCount("pizza") 
} // 실행시간 10초

func deliciousCount(food string) {
	for i := 0; i < 10; i++ {
		fmt.Println(food, "is delicious",i)
		time.Sleep(time.Second) 
	}
}