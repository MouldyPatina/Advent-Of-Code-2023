package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	racers := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(racers))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(racers))
}

func Part1Solver(racers []string) int {
	raceLen := getTimes(racers[0])
	bestTimes := getTimes(racers[1])
	if len(raceLen) != len(bestTimes) {
		log.Println("Error in inputs, but arrays should be the same length")
		return -1
	}
	output := 1
	for i := 0; i < len(raceLen); i++ {
		output *= numBetterTimes(raceLen[i], bestTimes[i])
	}
	return output
}

func Part2Solver(racers []string) int {
	raceLen := getSingleTimes(racers[0])
	bestTimes := getSingleTimes(racers[1])
	output := numBetterTimes(raceLen, bestTimes)
	return output
}

func numBetterTimes(leng int, time int) int {
	for i := 0; i < leng; i++ {
		if i*(leng-i) > time {
			return leng - 2*i + 1
		}
	}
	return -1
}

func getTimes(str string) []int {
	times := []int{}
	temp := 0
	pos := 0
	for str[pos] > '9' || str[pos] < '0' {
		pos++
	}
	for ; pos < len(str); pos++ {
		if digit := str[pos]; digit <= '9' && digit >= '0' {
			temp = temp*10 + int(digit) - int('0')
		} else if temp > 0 {
			times = append(times, temp)
			temp = 0
		}
	}
	times = append(times, temp)
	return times
}

func getSingleTimes(str string) int {
	temp := 0
	pos := 0
	for str[pos] > '9' || str[pos] < '0' {
		pos++
	}
	for ; pos < len(str); pos++ {
		if digit := str[pos]; digit <= '9' && digit >= '0' {
			temp = temp*10 + int(digit) - int('0')
		}
	}
	return temp
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
