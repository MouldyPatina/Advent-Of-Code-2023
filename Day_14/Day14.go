package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	platform := readFile("Test")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(platform))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(platform))
}

func Part1Solver(platform []string) int {
	var totalLoad = 0;
	for i := 0; i < len(platform[0]); i++ {
		rockLoad := len(platform);
		for j := 0; j < len(platform); j++ {
			switch (platform[j][i]) {
				case '#':
					rockLoad = len(platform) - j - 1;
				case 'O':
					totalLoad += rockLoad;
					rockLoad--;
			}
		}
	}
	return totalLoad;
}

func Part2Solver(platform []string) int {
	cyclePlatform := make([][]string, len(platform))
	for j := 0; j < len(cyclePlatform); j++ {
		cyclePlatform[j] = strings.Split(platform[j], "");
	}
	cyclePlatform = Cycled(cyclePlatform);
	var totalLoad = 0;
	for i := 0; i < len(cyclePlatform[0]); i++ {
		rockLoad := len(cyclePlatform);
		for j := 0; j < len(cyclePlatform); j++ {
			switch (cyclePlatform[j][i]) {
				case "#":
					rockLoad = len(cyclePlatform) - j - 1;
				case "O":
					totalLoad += rockLoad;
					rockLoad--;
			}
		}
	}
	return totalLoad;
}

func Cycled(platform [][]string) [][]string {
	rotatedPlatform := make([][]string, len(platform));
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
