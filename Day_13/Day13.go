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
	blockMirror := [][]string{};
	tempMirror := []string{};
	for _, mirror := range mirrors {
		if mirror != "" {
			tempMirror = append(tempMirror, mirror);
		} else {
			blockMirror = append(blockMirror, tempMirror);
			tempMirror = []string{};
		}
	}
	blockMirror = append(blockMirror, tempMirror);

	sum := 0;

	reflectionValue := 0;

	for _, block := range blockMirror {
		reflectionValue = HorizontalReflection(block) + VerticalReflection(block) * 100;
		sum += reflectionValue;
		reflectionValue = 0;
	}

	return sum;
}

func Part2Solver(mirrors []string) int {
	blockMirror := [][]string{};
	tempMirror := []string{};
	for _, mirror := range mirrors {
		if mirror != "" {
			tempMirror = append(tempMirror, mirror);
		} else {
			blockMirror = append(blockMirror, tempMirror);
			tempMirror = []string{};
		}
	}
	blockMirror = append(blockMirror, tempMirror);

	sum := 0;

	reflectionValue := 0;

	for _, block := range blockMirror {
		reflectionValue = HorizontalSmudge(block) + VerticalSmudge(block) * 100;
		sum += reflectionValue;
		reflectionValue = 0;
	}

	return sum;
}

func VerticalSmudge(pattern []string) int {
	for j := 1; j < len(pattern); j++ {
		reflected := true;
		smudged := false;
		size := Min(j, len(pattern) - j);
		for i := 0; i < len(pattern[j]) && reflected; i++ {
			for k := 0; k < size; k++ {
				reflected = reflected && (pattern[j - 1 - k][i] == pattern[j + k][i]);
				if !reflected && !smudged {
					smudged = true;
					reflected = true;
				}
			}
		}
		if reflected && smudged {
			return j;
		}
	}
	return 0;
}

func HorizontalSmudge(pattern []string) int {
	for i := 1; i < len(pattern[0]); i++ {
		reflected := true;
		smudged := false;
		size := Min(i, len(pattern[0]) - i);
		for j := 0; j < len(pattern) && reflected; j++ {
			for k := 0; k < size; k++ {
				reflected = reflected && (pattern[j][i - 1 - k] == pattern[j][i + k]);
				if !reflected && !smudged {
					smudged = true;
					reflected = true;
				}
			}
		}
		if reflected && smudged{
			return i;
		}
	}
	return 0;
}

func VerticalReflection(pattern []string) int {
	for j := 1; j < len(pattern); j++ {
		reflected := true;
		size := Min(j, len(pattern) - j);
		for i := 0; i < len(pattern[j]) && reflected; i++ {
			for k := 0; k < size; k++ {
				reflected = reflected && (pattern[j - 1 - k][i] == pattern[j + k][i]);			
			}
		}
		if reflected {
			return j;
		}
	}
	return 0;
}

func HorizontalReflection(pattern []string) int {
	for i := 1; i < len(pattern[0]); i++ {
		reflected := true;
		size := Min(i, len(pattern[0]) - i);
		for j := 0; j < len(pattern) && reflected; j++ {
			for k := 0; k < size; k++ {
				reflected = reflected && (pattern[j][i - 1 - k] == pattern[j][i + k]);			
			}
		}
		if reflected {
			return i;
		}
	}
	return 0;
}

func Min(a int, b int) int {
	if a < b {
		return a;
	}
	return b;
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
