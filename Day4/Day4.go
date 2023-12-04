package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func main() {
	calibrations := readFile("Real")
	fmt.Println(Part1Solver(calibrations))
	//fmt.Println(Part2Solver(calibrations));
}

func Part1Solver(scratchCard []string) int {
	sum := 0
	for j := 0; j < len(scratchCard); j++ {
		i := 0
		for scratchCard[j][i] != ':' {
			i++
		}
		i++
		winning := [100]bool{}
		score := 0
		var temp = 0
		for ; scratchCard[j][i] != '|'; i++ {
			if char := scratchCard[j][i]; char <= '9' && char >= '0' {
				temp = temp*10 + int(char) - int('0')
			} else if char == ' ' && temp != 0 {
				winning[temp] = true
				temp = 0
			}
		}
		i++
		for ; i < len(scratchCard[j]); i++ {
			if char := scratchCard[j][i]; char <= '9' && char >= '0' {
				temp = temp*10 + int(char) - int('0')
			} else if char == ' ' && temp != 0 {
				if winning[temp] {
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
				}
				temp = 0
			}
		}
		if winning[temp] {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
		sum += score
	}
	return sum
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
