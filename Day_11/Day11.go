package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	sky := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(sky))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(sky, 1000000))
}

func Part1Solver(sky []string) int {
	xExpansion := make([]int, len(sky[0]))
	yExpansion := make([]int, len(sky))
	galaxies := [][]int{}
	for j := 0; j < len(sky); j++ {
		for i := 0; i < len(sky[j]); i++ {
			if sky[j][i] == '#' {
				galaxies = append(galaxies, []int{j, i})
				xExpansion[i] = -1
				yExpansion[j] = -1
			}
		}
	}

	temp := 0
	for j := 0; j < len(yExpansion); j++ {
		yExpansion[j] += temp + 1
		temp = yExpansion[j]
	}
	temp = 0
	for i := 0; i < len(xExpansion); i++ {
		xExpansion[i] += temp + 1
		temp = xExpansion[i]
	}
	distance := 0

	for j := 0; j < len(galaxies); j++ {
		for k := j + 1; k < len(galaxies); k++ {
			xj, yj := galaxies[j][1], galaxies[j][0]
			xk, yk := galaxies[k][1], galaxies[k][0]
			if xk > xj {
				distance += xk - xj + xExpansion[xk] - xExpansion[xj]
			} else {
				distance += xj - xk + xExpansion[xj] - xExpansion[xk]
			}
			if yk > yj {
				distance += yk - yj + yExpansion[yk] - yExpansion[yj]
			} else {
				distance += yj - yk + yExpansion[yj] - yExpansion[yk]
			}
		}
	}
	return distance
}

func Part2Solver(sky []string, expansionRate int) int {
	xExpansion := make([]int, len(sky[0]))
	yExpansion := make([]int, len(sky))
	galaxies := [][]int{}
	expansionSize := expansionRate - 1
	for j := 0; j < len(sky); j++ {
		for i := 0; i < len(sky[j]); i++ {
			if sky[j][i] == '#' {
				galaxies = append(galaxies, []int{j, i})
				xExpansion[i] = -1 * expansionSize
				yExpansion[j] = -1 * expansionSize
			}
		}
	}

	temp := 0
	for j := 0; j < len(yExpansion); j++ {
		yExpansion[j] += temp + expansionSize
		temp = yExpansion[j]
	}
	temp = 0
	for i := 0; i < len(xExpansion); i++ {
		xExpansion[i] += temp + expansionSize
		temp = xExpansion[i]
	}
	distance := 0

	for j := 0; j < len(galaxies); j++ {
		for k := j + 1; k < len(galaxies); k++ {
			xj, yj := galaxies[j][1], galaxies[j][0]
			xk, yk := galaxies[k][1], galaxies[k][0]
			if xk > xj {
				distance += xk - xj + xExpansion[xk] - xExpansion[xj]
			} else {
				distance += xj - xk + xExpansion[xj] - xExpansion[xk]
			}
			if yk > yj {
				distance += yk - yj + yExpansion[yk] - yExpansion[yj]
			} else {
				distance += yj - yk + yExpansion[yj] - yExpansion[yk]
			}
		}
	}
	return distance
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
