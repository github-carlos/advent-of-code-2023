package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Almanac [][][]int;

func readInput(example bool) []string {

	var fileName string

	if example {
		fileName = "input_example.txt";
	} else {
		fileName = "input.txt";
	}

	input, _ := os.ReadFile(fileName);
	return strings.Split(string(input), "\n");
}

func main() {

	configs := make(Almanac, 7)

	input := removeEmptyLines(readInput(true))

	seeds := convArrayStrToInt(strings.Fields(strings.Split(input[0], ":")[1]));

	currentSlice := 1
	for index := range configs {
		extracted, i := extractConfigs(input[currentSlice:])
		configs[index] = extracted
		currentSlice += i + 1
	}


	for _, seed := range seeds {
		getLeastLocation(seed, 0, configs)
	}
}

func getLeastLocation(input int, position int, almanac Almanac) {

	result := calcFromAlmanac(input, almanac[position])

	if position == (len(almanac) - 1) {
		fmt.Println("Location", result);
		return;
	}

	position += 1
	getLeastLocation(result, position, almanac);
}

func calcFromAlmanac(input int, configs[][]int) int {

	result := input;
	for _, c := range configs {
		destinationStart := c[0]
		sourceStart := c[1]
		_range := sourceStart + c[2]

		if input > _range || input < sourceStart {
			continue;
		}

		result = destinationStart + (input - sourceStart)
		break;
	}

	return result;

}

func extractConfigs(almanac []string) ([][]int, int) {

	var configs [][]int;
	lastIndex := 0;
	for i, item := range almanac[1:] {
		splited := strings.Fields(item)
		if unicode.IsDigit(rune(splited[0][0])) {
			configs = append(configs, convArrayStrToInt(splited))
			continue;
		}
		lastIndex = i
		break;
	}
	return configs, lastIndex;
}

func convArrayStrToInt(list []string) (result []int) {

	for _, item := range list {
		n, _ := strconv.Atoi(item);
		result = append(result, n);
	}

	return;
}

func removeEmptyLines(arr []string) (filtered []string) {
	for _, i := range arr {
		if strings.TrimSpace(i) != "" {
			filtered = append(filtered, i)
		}
	}

	return filtered;
}

