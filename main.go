package main

import "fmt"

func main() {
	jihyun := map [string]string{"name": "jihyun", "age": "25"}
	for _, value := range jihyun {
		fmt.Println(value)
	}
}
