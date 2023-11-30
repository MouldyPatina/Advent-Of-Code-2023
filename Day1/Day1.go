package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Day 1, Advent of code 2022")
	file, err := os.Open("real.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var elf = []int{}
	var elfPack = 0
	for scanner.Scan() {
		if scanner.Text() != "" {
			val, err := strconv.ParseInt(scanner.Text(), 10, 0)
			if err == nil {
				elfPack += int(val)
			}
		} else {
			elf = append(elf, elfPack)
			elfPack = 0
		}
	}
	elf = append(elf, elfPack)
	fmt.Println(Max(elf))
	var part2 = Top3(elf)
	fmt.Println(part2[0] + part2[1] + part2[2])
}

func Max(input []int) int {
	var max int = 0
	for _, ele := range input {
		if ele > max {
			max = ele
		}
	}
	return max
}

func Top3(input []int) []int {
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})
	return []int{input[0], input[1], input[2]}
}
