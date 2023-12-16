package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	sequence := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(sequence))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(sequence))
}

func Part1Solver(ascii []string) int {
	hashSum := 0;
	for _, word := range strings.Split(ascii[0], ",") {
		hashValue := 0;
		for i := 0; i < len(word); i++ {
			hashValue = ((hashValue + (int)(word[i])) * 17) % 256;
		}
		hashSum += hashValue;
		hashValue = 0;
	}
	return hashSum;
}

func Part2Solver(ascii []string) int {
	return -1;
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
