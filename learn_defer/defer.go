package main

import (
	"fmt"
)

func runLater() {
	fmt.Println("should come later")
}

func runFirst() {
	fmt.Println("should come first")
}

func runLast() {
	fmt.Println("this should run last")
}

func run() {
	defer runLast()
	defer runLater()
	runFirst()
}

func main() {
	fmt.Println("learn how to use defer in golang")

	run()
}
