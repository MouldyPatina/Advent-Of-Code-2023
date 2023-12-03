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
	fmt.Println(Part2Solver(calibrations));
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

func Part2Solver(engine []string) int {
	gears := 0;
	for j := 0; j < len(engine); j++ {
		for i := 0; i < len(engine[j]); i++ {
			if engine[j][i] == '*' {
				gear := GetGears(i, j, engine);
				log.Println("Gear is at x:", i, "y:", j, " len is:", len(gear));
				if len(gear) == 2 {
					log.Println("Gear is", gear[0], gear[1]);
					gears += gear[0] * gear[1];
				}
			}
		}
	}
	return gears;
}

func GetGears(x int, y int, engine []string) []int {
	gear := []int{};
	for j := Max(0, y - 1); j < Min(len(engine), y + 2); j++ {
		skip := false;
		for i := Max(0, x - 1); i < Min(len(engine[j]), x + 2); i++ {
			if letter := engine[j][i]; letter <= '9' && letter >= '0' {
				if !skip {
					gear = append(gear, GetNum(engine[j], i));
				}
				skip = true;
			} else {
				skip = false;
			}
		}
	}
	return gear;
}

func GetNum(str string, pos int) int {
	for pos > -1 && str[pos] <= '9' && str[pos] >= '0' {
		pos--;
	}
	pos = Max(0, pos + 1);
	temp := 0;
	for ;pos < len(str) && str[pos] <= '9' && str[pos] >= '0'; pos++ {
		temp = temp * 10 + int(str[pos]) - int('0');
	}
	return temp;
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
