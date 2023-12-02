package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("Test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileContent := bufio.NewScanner(file)
	var games = []string{}
	for fileContent.Scan() {
		games = append(games, fileContent.Text())
	}

	fmt.Println("Solution to Part 1")
	var possible = PossiblePart1(games)
	fmt.Println(Sum(possible));

	fmt.Println("Solution to Part 2")
	var powers = PowerPart2(calibrations)
	fmt.Println(Sum(powers));
}

func PossiblePart1(games []string) []int {
	var possible = []int;
	for j := 0; j < len(games); j++ {
		i := 0;
		for ; games[j][i] != ':'; i++ {}
		i++;
		var temp = 0;
		for ; i < len(games[j]); i++ {
			if word := games[j][i]; word <= '9' && word >= '0' {
				temp = temp * 10 + int(word) - int('0');
			}
		}
		if i == len(games) {
			possible = append(possible, j + 1);
		}
	}
}

func PowerPart2(games []string) []int {

}

func Sum(input []int) int {
	output := 0
	for _, val := range input {
		output += val;
	}
	return output;
}
