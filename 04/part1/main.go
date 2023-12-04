package main

import (
	"fmt"
	"math/rand"
)

func main() {
// 	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	// splitedCards := strings.Split(input, "\n")
	fmt.Println(quickSort([]int{2, 3, 5, 9, 1, 8, 3, 19, 20, 12}))
	// fmt.Println(quickSort([]int{2, 3, 5, 9, 1, 8, 3}))
}

func quickSort(numbers[]int) []int {

	if len(numbers) < 2 {
		return numbers;
	}
	randIndex := rand.Intn(len(numbers))
	pivot := numbers[randIndex]

	var left []int
	var right []int

	for i, item := range numbers {

		if i == randIndex {
			continue;
		}

		if item <= pivot {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}

	fmt.Println("Pivot", pivot)
	fmt.Println("left", left, "right", right)
	left = append(quickSort(left), pivot)
	return append(left, quickSort(right)...);
}

func binarySearch(search int, numbers []int) (position int) {
	return 1;
}
