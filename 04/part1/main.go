package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	input := string(file)

	cards := strings.Split(input, "\n")

	result := 0

	multipliers := make(map[int]int)

	for i, card := range cards {
		if card == "" {
			continue
		}
		splitedCard := strings.Split(card, " | ")
		winNumbersStr := strings.Split(splitedCard[0], " ")[2:]

		var winNumbers []int
		for _, n := range winNumbersStr {
			winNumbers = append(winNumbers, convToInt(n))
		}

		playedNumbersStr := strings.Split(splitedCard[1], " ")

		var playedNumbers []int
		for _, n := range playedNumbersStr {
			if n == "" {
				continue
			}
			playedNumbers = append(playedNumbers, convToInt(n))
		}
		playedNumbers = quickSort(playedNumbers)

		fmt.Println("Winners", winNumbers)
		fmt.Println("Played", playedNumbers)

		points := 0
		propagateMultiplier := 0
		for _, w := range winNumbers {
			_, err := binarySearch(w, playedNumbers)
			if err == nil {
				propagateMultiplier++
				if points == 0 {
					points++
				} else {
					points *= 2
				}
			}
		}
		multiplier := multipliers[i] + 1
		fmt.Println("i", i)
		fmt.Println("Multiplier", multiplier)
		fmt.Println("Points:", points*multiplier)
		result += multiplier

		fmt.Println("Propagate", propagateMultiplier)
		for j := i + 1; j <= i+propagateMultiplier; j++ {
			multipliers[j] = multipliers[j] + multiplier
		}
		fmt.Println("Map", multipliers)
	}
	fmt.Println("Final Result:", result)
}

func quickSort(numbers []int) []int {

	if len(numbers) < 2 {
		return numbers
	}
	randIndex := rand.Intn(len(numbers))
	pivot := numbers[randIndex]

	var left []int
	var right []int

	for i, item := range numbers {

		if i == randIndex {
			continue
		}

		if item <= pivot {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}

	left = append(quickSort(left), pivot)
	return append(left, quickSort(right)...)
}

func binarySearch(search int, numbers []int) (position int, err error) {
	left := 0
	right := len(numbers)

	var middlePosition int

	for {
		arr := numbers[left:right]
		if len(arr) == 0 {
			err = errors.New("Not found")
			break
		}

		middlePosition = (right - left) / 2
		value := arr[middlePosition]
		if value == search {
			position = left + middlePosition
			break
		}

		if middlePosition == 0 {
			middlePosition = 1
		}

		if search < value {
			right = right - middlePosition
		} else {
			left = left + middlePosition
		}
	}

	return position, err
}

func convToInt(v string) int {
	converted, _ := strconv.Atoi(v)
	return converted
}
