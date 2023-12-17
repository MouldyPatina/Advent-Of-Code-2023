package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	mirrors := readFile("Test")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(mirrors))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(mirrors))
}

func Part1Solver(mirrors []string) int {
	laser := make([][]int, len(mirrors));
	for i := 0; i < len(mirrors); i++ {
		laser[i] = make([]int, len(mirrors[i]));
	}
	FireLaser(mirrors, laser, 'R', []int{0,0});
	return Sum(laser);
}

func FireLaser(mirror []string, light [][]int, move byte, pos []int) {
	if y, x := pos[0], pos[1]; x < 0 || x >= len(mirror[0]) || y < 0 || y >= len(mirror) {
		return;
	}
	light[pos[0]][pos[1]] = 1;
	if mir := mirror[pos[0]][pos[1]]; ((mir == '|' || mir == '.') && move == 'U') || (mir == '\\' && move == 'L') || (mir == '/' && move == 'R') {
		FireLaser(mirror, light, 'U', []int{pos[0] - 1, pos[1]});
	} else if mir := mirror[pos[0]][pos[1]]; ((mir == '|' || mir == '.') && move == 'D') || (mir == '\\' && move == 'R') || (mir == '/' && move == 'L') {
		FireLaser(mirror, light, 'D', []int{pos[0] + 1, pos[1]});
	} else if mir := mirror[pos[0]][pos[1]]; ((mir == '-' || mir == '.') && move == 'L') || (mir == '\\' && move == 'U') || (mir == '/' && move == 'D') { 
		FireLaser(mirror, light, 'L', []int{pos[0], pos[1] - 1});
	} else if mir := mirror[pos[0]][pos[1]]; ((mir == '-' || mir == '.') && move == 'R') || (mir == '\\' && move == 'D') || (mir == '/' && move == 'U') {
		FireLaser(mirror, light, 'R', []int{pos[0], pos[1] + 1});
	} else if mir := mirror[pos[0]][pos[1]]; mir == '-' && (move == 'U' || move == 'D') {
		FireLaser(mirror, light, 'R', []int{pos[0], pos[1] + 1});
		FireLaser(mirror, light, 'L', []int{pos[0], pos[1] - 1});
	} else if mir := mirror[pos[0]][pos[1]]; mir == '|' && (move == 'L' || move == 'R') {
		FireLaser(mirror, light, 'D', []int{pos[0] + 1, pos[1]});
		FireLaser(mirror, light, 'U', []int{pos[0] - 1, pos[1]});
	}
}

func Sum(matrix [][]int) int {
	sum := 0;
	for _, line := range matrix {
		for _, num := range line {
			sum += num;
		}
	}
	return sum;
}

func Part2Solver(mirrors []string) int {
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
