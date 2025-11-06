package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}
func main() {
	favFood := []string{"kimchi, ramen"}
	jihyun := person{name: "jihyun", age: 25, favFood: favFood}
	fmt.Println(jihyun)
}
