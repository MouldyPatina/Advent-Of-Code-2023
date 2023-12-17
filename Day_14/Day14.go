package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	platform := readFile("Real")
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
	cache := make(map[string] int);	
	cache[ToString(cyclePlatform)] = 0;
	for k := 0; k < 1000000000; k++ {
		Cycled(cyclePlatform);
		if cache[ToString(cyclePlatform)] == 0 {
			cache[ToString(cyclePlatform)] = k;
		} else {
			first := cache[ToString(cyclePlatform)];
			for t := 0; t < (1000000000 - first - 1) % (k - first); t++ {
				Cycled(cyclePlatform);
			}
			break;
		}
	}
	var totalLoad = 0;
	for i := 0; i < len(cyclePlatform[0]); i++ {
		rockLoad := len(cyclePlatform);
		for j := 0; j < len(cyclePlatform); j++ {
			switch (cyclePlatform[j][i]) {
				case "#":
					rockLoad = len(cyclePlatform) - j - 1;
				case "O":
					totalLoad += len(cyclePlatform) - j;
					rockLoad--;
			}
		}
	}
	return totalLoad;
}

func ToString(matrix [][]string) string {
	var output = "";
	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(matrix[j]); i++ {
			output += matrix[j][i];
		}
	}
	return output;
}

func Cycled(platform [][]string) {
	TiltUp(platform);
	TiltRight(platform);
	TiltDown(platform);
	TiltLeft(platform);
	temp := "";
	for _, line := range platform {
		for _, rock:= range line {
			temp += rock
		}
		temp += "\n";
	}
}

func TiltUp(platform [][]string) {
	for i := 0; i < len(platform[0]); i++ {
		rollUp := 0;
		for j := 0; j < len(platform); j++ {
			switch platform[j][i] {
			case "#":
				rollUp = j + 1;
			case "O":
				platform[j][i] = ".";
				platform[rollUp][i] = "O";
				rollUp++;
			}
		}
	}
}

func TiltRight(platform [][]string) {
	for j := 0; j < len(platform); j++ {
		rollUp := 0;
		for i := 0; i < len(platform[j]); i++ {
			switch platform[j][i] {
			case "#":
				rollUp = i + 1;
			case "O":
				platform[j][i] = ".";
				platform[j][rollUp] = "O";
				rollUp++;
			}
		}
	}
}

func TiltDown(platform [][]string) {
	for i := 0; i < len(platform[0]); i++ {
		rollUp := len(platform) - 1;
		for j := len(platform) - 1; j > -1; j-- {
			switch platform[j][i] {
			case "#":
				rollUp = j - 1;
			case "O":
				platform[j][i] = ".";
				platform[rollUp][i] = "O";
				rollUp--;
			}
		}
	}
}

func TiltLeft(platform [][]string) {
	for j := 0; j < len(platform); j++ {
		rollUp := len(platform[j]) - 1;
		for i := len(platform[j]) - 1; i > -1; i-- {
			switch platform[j][i] {
			case "#":
				rollUp = i - 1;
			case "O":
				platform[j][i] = ".";
				platform[j][rollUp] = "O";
				rollUp--;
			}
		}
	}
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
