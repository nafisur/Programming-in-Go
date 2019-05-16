package main

import (
	"fmt"
	"log"
)

func main() {
	var number float64

	fmt.Printf("enter a float value\n")
	_, err := fmt.Scanln(&number)
	if err != nil {
		log.Fatalf("Failed to read because %s", err)
	}

	fmt.Printf("%d\n", int64(number))
}
