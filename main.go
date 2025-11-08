package main

import (
	"fmt"
	"time"
)

func main() {
	// make channel
	c := make(chan string)
	people := [5]string{"jihyun", "gildong", "sam", "john", "anna"}

	for _, person := range people {
		go isFriend(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for", i)
		fmt.Println(<- c)
	}
}

// isFriend: 5초 뒤 true라는 메시지를 보내줌
func isFriend(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is my friend"
}
