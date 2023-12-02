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
	var games = []string{}
	for fileContent.Scan() {
		games = append(games, fileContent.Text())
	}

	fmt.Println("Solution to Part 1")
	var possible = PossiblePart1(games)
	fmt.Println(Sum(possible));

	fmt.Println("Solution to Part 2")
	var powers = PowerPart2(games)
	fmt.Println(Sum(powers));
}

func PossiblePart1(games []string) []int {
	red := 12;
	green := 13;
	blue := 14;
	var possible = []int{};
	for j := 0; j < len(games); j++ {
		i := 0;
		for ; games[j][i] != ':'; i++ {}
		i++;
		var temp = 0;
		for ; i < len(games[j]); i++ {
			if word := games[j][i]; word <= '9' && word >= '0' {
				temp = temp * 10 + int(word) - int('0');
			} else if word == 'r' {
				if temp > red {
					possible = append(possible, 0);
					break;
				}
				temp = 0;
			} else if word == 'g' {
				if temp > green {
					possible = append(possible, 0);
					break;
				}
				temp = 0;
			} else if word == 'b' {
				if temp > blue {
					possible = append(possible, 0);
					break;
				}
				temp = 0;
			}
		}
		if i == len(games[j]) {
			possible = append(possible, j + 1);
		}
	}
	return possible;	
}

func PowerPart2(games []string) []int {
var powers = []int{};
	for j := 0; j < len(games); j++ {
		i := 0;
		red := 0;
		green := 0;
		blue := 0;
	
		for ; games[j][i] != ':'; i++ {}
		i++;
		var temp = 0;
		for ; i < len(games[j]); i++ {
			if word := games[j][i]; word <= '9' && word >= '0' {
				temp = temp * 10 + int(word) - int('0');
			} else if word == 'r' {
				red = max(red, temp);	
				temp = 0;
			} else if word == 'g' {
				green = max(green, temp);	
				temp = 0;
			} else if word == 'b' {
				blue = max(blue, temp);	
				temp = 0;
			}
		}
		powers = append(powers, blue * green * red);
	}
	return powers;	
}

func Sum(input []int) int {
	output := 0
	for _, val := range input {
		output += val;
	}
	return output;
}

func max(a int, b int) int {
	if a > b {
		return a;
	}
	return b;
}
