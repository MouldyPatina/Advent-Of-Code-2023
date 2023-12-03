package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
)

func init () {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file);
}

func main() {
	calibrations := readFile("Real");
	fmt.Println(Part1Solver(calibrations));
}

func Part1Solver(calibration []string) int {
	sum := 0;
	for j := 0; j < len(calibration); j++ {
		x := [2]int{-1, -1};
		y := j;
		temp := 0;
		for i := 0; i < len(calibration[j]); i++ {
			if letter := calibration[j][i]; letter >= '0' && letter <= '9' {
				temp = temp * 10 + int(letter) - int('0');
				if x[0] == -1 {
					x[0] = i;
					x[1] = i;
				} else {
					x[1] = i;
				}
			} else if x[0] != -1 && x[1] != -1 {
				if NearSymbol(x, y, calibration) {
					sum += temp;
				}
				x = [2]int{-1, -1};
				temp = 0;
			}
		}
		if x[0] != -1 && x[1] != -1 {
			if NearSymbol(x, y, calibration) {
				sum += temp;
			}
		}
	}
	return sum;
}

func NearSymbol(x [2]int, y int, calibration []string) bool {
	for j := Max(0, y - 1); j < Min(len(calibration), y + 2); j++ {
		for i := Max(0, x[0] - 1); i < Min(len(calibration[j]), x[1] + 2); i++ {
			if j == y && i >= x[0] && i <= x[1] {
			} else if word := calibration[j][i]; (word > '9' || word < '0') && word != '.' {
				return true;
			}
		}
	}
	return false;
}

func Max(a int, b int) int {
	if a > b {
		return a;
	}
	return b;
}

func Min(a int, b int) int {
	if a < b {
		return a;
	}
	return b;
}

func readFile(name string) []string {
	file, err := os.Open(name + ".txt");
	if err != nil {
		log.Println("Failed to read file as:", err);
		panic(err);
	}
	defer file.Close();
	fileContent := bufio.NewScanner(file);
	content := []string{};
	for fileContent.Scan() {
		content = append(content, fileContent.Text());
	}
	return content;
}
