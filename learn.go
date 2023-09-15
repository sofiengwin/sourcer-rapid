package main

import (
	"fmt"
)

type base struct {
	num int
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	symbol := [...]string{USD: "$", EUR: "#", GBP: "##", 4: "5", 100: "100"}

	for key, value := range symbol {
		fmt.Printf("key: %d Value: %v\n", key, value)
	}
	fmt.Printf("Hello World! %v\n", USD)
}
