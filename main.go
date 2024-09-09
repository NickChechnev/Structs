package main

import (
	"fmt"
	"structs/magazine"
)

func main() {
	s := new(magazine.Subscriber)
	s.Rate = 4.99
	fmt.Println(s.Rate)
}
