package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	names = append(names, "flynn")
	fmt.Println(names)
}
