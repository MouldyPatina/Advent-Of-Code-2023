package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	mirrors := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(mirrors))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(mirrors))
}

func Part1Solver(mirrors []string) int {
	laser := make([][][]byte, len(mirrors));
	for i := 0; i < len(mirrors); i++ {
		laser[i] = make([][]byte, len(mirrors[i]));
	}
	FireLaser(mirrors, laser, 'R', []int{0,0});
	return Sum(laser);
}

func Part2Solver(mirrors []string) int {
	laser := make([][][]byte, len(mirrors));
	for i := 0; i < len(mirrors); i++ {
		laser[i] = make([][]byte, len(mirrors[i]));
	}
	maxEnergy := 0;
	for i := 0; i < len(mirrors); i++ {
		FireLaser(mirrors, laser, 'R', []int{i,0});
		tempEnergy := Sum(laser);
		maxEnergy = Max(tempEnergy, maxEnergy);
		MakeEmpty(laser);
		FireLaser(mirrors, laser, 'R', []int{i,0})
		tempEnergy = Sum(laser);
		maxEnergy = Max(tempEnergy, maxEnergy);
		MakeEmpty(laser);
	}
	for i := 0; i < len(mirrors[0]); i++ {
		FireLaser(mirrors, laser, 'D', []int{0,i})
		tempEnergy := Sum(laser);
		maxEnergy = Max(tempEnergy, maxEnergy);
		MakeEmpty(laser);
		FireLaser(mirrors,laser,'U',[]int{len(mirrors)-1,i})
		tempEnergy = Sum(laser);
		maxEnergy = Max(tempEnergy, maxEnergy);
		MakeEmpty(laser);
	}
	return maxEnergy;
}

func MakeEmpty(matrix [][][]byte) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = []byte{};
		}
	}
}

func Max(a int, b int) int {
	if a > b {
		return a;
	}
	return b;
}

func FireLaser(mirror []string, light [][][]byte, move byte, pos []int) {
	if y, x := pos[0], pos[1]; x < 0 || x >= len(mirror[0]) || y < 0 || y >= len(mirror) {
		return;
	}
	if len(light[pos[0]][pos[1]]) > 0 {
		for _, moved := range light[pos[0]][pos[1]] {
			if moved == move {
				return;
			}
		}
	}
	light[pos[0]][pos[1]] = append(light[pos[0]][pos[1]], move);
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

func Sum(matrix [][][]byte) int {
	sum := 0;
	for _, line := range matrix {
		for _, num := range line {
			if len(num) > 0 {
				sum++;
			}
		}
	}
	return sum;
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
