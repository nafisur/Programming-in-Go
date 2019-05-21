package main

import (
	"fmt"
	"log"
)

const count = 10

func main() {

	inputn := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Enter an integer %d: ", i)
		if _, err := fmt.Scanf("%d", &inputn[i]); err != nil {
			log.Fatalf("Invalid input\n")
		}
	}

	BubbleSort(inputn)
	fmt.Printf("Sorted array: %d\n", inputn)

}

func BubbleSort(inputn []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(inputn)-1; i++ {
			if inputn[i+1] < inputn[i] {
				Swap(inputn, i)
				swapped = true
			}
		}
	}
}

func Swap(inputn []int, current int) {
	inputn[current], inputn[current+1] = inputn[current+1], inputn[current]
}
