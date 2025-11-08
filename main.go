package main

import (
	"fmt"
	"time"
)

func main() {
	// make channel
	c := make(chan bool)

	people := [2]string{"jihyun", "gildong"}
	for _, person := range people {
		go isFriend(person, c)
	}
	fmt.Println(<- c)
	fmt.Println(<- c)
	fmt.Println(<- c)
}

// isFriend: 5초 뒤 true라는 메시지를 보내줌
func isFriend(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
