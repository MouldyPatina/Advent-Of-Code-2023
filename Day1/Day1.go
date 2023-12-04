package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("Real.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileContent := bufio.NewScanner(file)
	var calibrations = []string{}
	for fileContent.Scan() {
		calibrations = append(calibrations, fileContent.Text())
	}

	fmt.Println("Solution to Part 1")
	var values = realCalibrationPart1(calibrations)
	fmt.Println(Sum(values))

	fmt.Println("Solution to Part 2")
	var values2 = realCalibrationPart2(calibrations)
	fmt.Println(Sum(values2))
}

func realCalibrationPart1(calib []string) []int {
	var output = []int{}
	for _, line := range calib {
		val := 0
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				val = int(line[i]) - int('0')
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				val = val*10 + int(line[i]) - int('0')
				break
			}
		}
		output = append(output, val)
	}
	return output
}

func realCalibrationPart2(calib []string) []int {
	var output = []int{}
	for _, line := range calib {
		val := 0
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				val = int(line[i]) - int('0')
				break
			}
			if WrittenNumber(i, i+3, line, "one") {
				val = 1
				break
			} else if WrittenNumber(i, i+3, line, "two") {
				val = 2
				break
			} else if WrittenNumber(i, i+3, line, "six") {
				val = 6
				break
			} else if WrittenNumber(i, i+4, line, "four") {
				val = 4
				break
			} else if WrittenNumber(i, i+4, line, "five") {
				val = 5
				break
			} else if WrittenNumber(i, i+4, line, "nine") {
				val = 9
				break
			} else if WrittenNumber(i, i+5, line, "three") {
				val = 3
				break
			} else if WrittenNumber(i, i+5, line, "seven") {
				val = 7
				break
			} else if WrittenNumber(i, i+5, line, "eight") {
				val = 8
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				val = val*10 + int(line[i]) - int('0')
				break
			}
			if WrittenNumber(i-2, i+1, line, "one") {
				val = val*10 + 1
				break
			} else if WrittenNumber(i-2, i+1, line, "two") {
				val = val*10 + 2
				break
			} else if WrittenNumber(i-2, i+1, line, "six") {
				val = val*10 + 6
				break
			} else if WrittenNumber(i-3, i+1, line, "four") {
				val = val*10 + 4
				break
			} else if WrittenNumber(i-3, i+1, line, "five") {
				val = val*10 + 5
				break
			} else if WrittenNumber(i-3, i+1, line, "nine") {
				val = val*10 + 9
				break
			} else if WrittenNumber(i-4, i+1, line, "three") {
				val = val*10 + 3
				break
			} else if WrittenNumber(i-4, i+1, line, "seven") {
				val = val*10 + 7
				break
			} else if WrittenNumber(i-4, i+1, line, "eight") {
				val = val*10 + 8
				break
			}
		}
		output = append(output, val)
	}
	return output
}

func WrittenNumber(start int, end int, str string, number string) bool {
	if start < 0 || end > len(str) {
		return false
	} else if str[start:end] == number {
		return true
	}
	return false
}

func Sum(input []int) int {
	output := 0
	for _, val := range input {
		output += val
	}
	return output
}
