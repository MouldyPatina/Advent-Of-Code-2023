package main 

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	file := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(file))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(file))
}

func Part1Solver(oasis []string) int {
	sequences := Sequence(oasis);

	sumNewSeq := 0;
	for _, sequence := range sequences {
		sumNewSeq += PredictElement(sequence, true);
	}
	
	return sumNewSeq;
}

func Part2Solver(oasis []string) int {
	sequences := Sequence(oasis);

	sumNewSeq := 0;
	for _, sequence := range sequences {
		sumNewSeq += PredictElement(sequence, false);
	}
	
	return sumNewSeq;
}

//If NextEle true will predict element at end of sequence, else will predict at the start
func PredictElement(sequence []int, NextEle bool) int {
	numElem := len(sequence);
	diffSeq := make([]int, numElem - 1);
	diffZero := true;
	for i := 0; i < numElem - 1; i++ {
		diffSeq[i] = sequence[i + 1] - sequence[i];
		diffZero = diffZero && diffSeq[i] == 0;
	}
	if diffZero {
		return sequence[0];
	}
	
	findLast := 1
	if !NextEle {
		findLast = -1;
		numElem = 1;
	}
	return sequence[numElem - 1] + findLast * PredictElement(diffSeq, NextEle);
}
	

func Sequence(strList []string) [][]int {
	output := make([][]int, len(strList))
	for i := 0; i < len(output); i++ {
		tempSequence := []int{};
		for _, num := range strings.Split(strList[i], " ") {
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Println("Couldn't parse int", err);
				break;
			}
			tempSequence = append(tempSequence, val);
		}
		output[i] = tempSequence;
	}
	return output;
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
