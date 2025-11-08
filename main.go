package main

import (
	"fmt"
	"time"
)

func main() {
	go countPerson("jihyun")
	go countPerson("gildong")
	time.Sleep(time.Second * 5)
}

func countPerson(person string) {
	for i := 0;i < 10; i++ {
		fmt.Println(person, i)
		time.Sleep(time.Second)
	}
}