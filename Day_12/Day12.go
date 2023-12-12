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
	springsMap := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(springsMap))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(springsMap))
}

func Part1Solver(springs []string) int {
	conditions := []Spring{};

	for _, spring := range springs {
		temp := strings.Split(spring, " ") 
		tempSpring := Spring{};
		tempSpring.list = strings.Split(temp[0], "")
		tempInt := []int{};
		for _, num := range strings.Split(temp[1], ",") {
			i, err := strconv.Atoi(num);
			if err != nil {
				log.Println("Couldn't parse int for damaged strings as:", err)
			}
			tempInt = append(tempInt, i);
		}
		tempSpring.numDam = tempInt;
		conditions = append(conditions, tempSpring)
	}

	combinations := 0;

	for _, con := range conditions {
		combinations += GetCombinations(con.list, con.numDam)
	}
	return combinations;
}

func Part2Solver(springs []string) int {
	conditions := []Spring{};

	for _, spring := range springs {
		temp := strings.Split(spring, " ") 
		tempSpring := Spring{};
		tempSpring.list = strings.Split(temp[0], "")
		tempInt := []int{};
		for _, num := range strings.Split(temp[1], ",") {
			i, err := strconv.Atoi(num);
			if err != nil {
				log.Println("Couldn't parse int for damaged strings as:", err)
			}
			tempInt = append(tempInt, i);
		}
		tempSpring.numDam = tempInt;
		conditions = append(conditions, tempSpring)
	}

	combinations := 0;
	i := 0;
	for _, con := range conditions {
		tempList := con.list;
		tempDam := con.numDam; 

		for i := 0; i < 4; i++ {
			tempList = append(tempList, "?");
			tempList = append(tempList, con.list...);
			tempDam = append(tempDam, con.numDam...);
		}
		fmt.Println(i + 1)
		i++
		val := GetCombinations(tempList, tempDam)
		combinations += val
	}
	return combinations;
}

func GetCombinations(cond []string, numDam []int) int {
	pos := 0;
	brokeSum := 0;
	for i := 0; i < len(cond); i++ {
		if cond[i] == "?" {
			copyString := make([]string, len(cond))
			copy(copyString, cond)
			copyString[i] = "#"
			sum := GetCombinations(copyString, numDam)
			copy(copyString, cond)
			copyString[i] = "."
			sum += GetCombinations(copyString, numDam)
			return sum;
		} else if cond[i] == "." && brokeSum != 0 {
			if !(pos < len(numDam) && numDam[pos] == brokeSum) {
				return 0;
			} 
			brokeSum = 0
			pos++
		} else if cond[i] == "#" {
			brokeSum++
			if pos >= len(numDam) || numDam[pos] < brokeSum {
				return 0;
			}
		}
	}
	if (pos == len(numDam) - 1 && numDam[pos] == brokeSum) || pos == len(numDam) {
		return 1;
	}

	return 0;
}

type Spring struct {
	list []string
	numDam []int
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
