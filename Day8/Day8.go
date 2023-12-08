package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	direction := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(direction))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(direction))
}

func Part1Solver(direction []string) int {
	moves := direction[0]
	location := make(map[string][2]string)
	for _, directs := range direction[2:] {
		temp := GetStrings(directs)
		location[temp[0]] = [2]string{temp[1], temp[2]}
	}

	start := "AAA"
	Journey := 0

	for ; start != "ZZZ"; Journey++ {
		if moves[Journey%len(moves)] == 'L' {
			start = location[start][0]
		} else if moves[Journey%len(moves)] == 'R' {
			start = location[start][1]
		} else {
			log.Println("invalid direction")
		}
	}

	return Journey
}

func Part2Solver(direction []string) int {
	moves := direction[0]
	location := make(map[string][2]string)
	start := []string{}
	for _, directs := range direction[2:] {
		temp := GetStrings(directs)
		location[temp[0]] = [2]string{temp[1], temp[2]}
		if temp[0][2] == 'A' {
			start = append(start, temp[0])
		}
	}

	Journey := make([]int, len(start))

	for i := 0; i < len(start); i++ {
		for ; start[i][2] != 'Z'; Journey[i]++ {
			if moves[Journey[i]%len(moves)] == 'L' {
				start[i] = location[start[i]][0]
			} else if moves[Journey[i]%len(moves)] == 'R' {
				start[i] = location[start[i]][1]
			} else {
				log.Println("invalid direction")
			}
		}
	}

	for _, joun := range Journey {
		fmt.Println(joun)
	}
	return -1
}

func GetStrings(str string) []string {
	output := []string{}
	pos := 0
	for i := 0; i < len(str); i++ {
		if (str[i] > 'Z' || str[i] < 'A') && (str[i] > '9' || str[i] < '0') {
			if pos < i {
				output = append(output, str[pos:i])
			}
			pos = i + 1
		}
	}
	output = append(output, str[pos:])
	return output
}

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func readFile(name string) []string {
	file, err := os.Open(name + ".txt")
	if err != nil {
		log.Println("Failed to read file as:", err)
		panic(err)
	}
	defer file.Close()
	fileContent := bufio.NewScanner(file)
	content := []string{}
	for fileContent.Scan() {
		content = append(content, fileContent.Text())
	}
	return content
}
