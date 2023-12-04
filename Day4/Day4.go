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
	scratchCards := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(scratchCards))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Sum(Part2Solver(scratchCards)))
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
					score++
				}
				temp = 0
			}
		}
		if winning[temp] {
			score++
		}
		if score != 0 {
			sum += 1 << (score - 1)
		}
	}
	return sum
}

func Part2Solver(scratchCard []string) []int {
	wonScoreCards := []int{}
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
					score += 1
				}
				temp = 0
			}
		}
		if winning[temp] {
			score++
		}

		if len(wonScoreCards) <= j {
			wonScoreCards = append(wonScoreCards, 1)
		} else {
			wonScoreCards[j]++
		}

		for cards := j + 1; cards < j+score+1; cards++ {
			if len(wonScoreCards) <= cards {
				wonScoreCards = append(wonScoreCards, wonScoreCards[j])
			} else {
				wonScoreCards[cards] += wonScoreCards[j]
			}
		}
	}
	return wonScoreCards
}

func Sum(input []int) int {
	output := 0
	for _, number := range input {
		output += number
	}
	return output
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
