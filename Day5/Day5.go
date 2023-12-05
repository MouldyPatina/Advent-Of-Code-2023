package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	almanac := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(almanac))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(almanac))
}

func Part1Solver(almanac []string) uint32 {
	seeds := GetUInts32(almanac[0])
	maps := GetMaps(almanac[2:])
	for _, mp := range maps {
		seeds = convertSeeds(seeds, mp)
	}
	return min(seeds)
}

func Part2Solver(almanac []string) uint32 {
	seeds := GetUInts32(almanac[0])
	maps := GetMaps(almanac[2:])
	for _, mp := range maps {
		seeds = convertSeedsPart2(seeds, mp)
	}
	return minPart2(seeds)
}

type AlmanacMap struct {
	name        string
	Convertions []Convert
}

type Convert struct {
	destination uint32
	source      uint32
	rangeLength uint32
}

func convertSeedsPart2(seeds []uint32, convert AlmanacMap) []uint32 {
	changedSeeds := []uint32{}
	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		lenSeed := seeds[i+1]
		for _, update := range convert.Convertions {
			if seed >= update.source && seed < update.rangeLength+update.source {
				changedSeeds = append(changedSeeds,
					seed+update.destination-update.source,
					minNum(update.rangeLength, lenSeed))
				seed = seed + update.rangeLength
				if update.rangeLength >= lenSeed {
					lenSeed = 0
					break
				}
				lenSeed = lenSeed - update.rangeLength
			} else if seed < update.source && lenSeed+seed >= update.source {
				changedSeeds = append(changedSeeds,
					update.destination,
					lenSeed-update.source+seed)
				if lenSeed+seed > update.source+update.rangeLength {
					log.Println("did not handle this case rip")
				}
				lenSeed = update.source - seed
			}
		}
		if lenSeed != 0 {
			changedSeeds = append(changedSeeds, seed, lenSeed)
		}
	}
	return changedSeeds
}

func convertSeeds(seeds []uint32, convert AlmanacMap) []uint32 {
	for i := 0; i < len(seeds); i++ {
		seed := seeds[i]
		for _, update := range convert.Convertions {
			if seed >= update.source && seed < update.rangeLength+update.source {
				seed = seed + update.destination - update.source
				break
			}
		}
		seeds[i] = seed
	}
	return seeds
}

func GetMaps(almanac []string) []AlmanacMap {
	maps := []AlmanacMap{}

	tempMap := AlmanacMap{}

	for _, line := range almanac {
		if line == "" {
			maps = append(maps, tempMap)
			tempMap = AlmanacMap{}
		} else if line[0] <= '9' && line[0] >= '0' {
			nums := GetUInts32(line)
			tempMap.Convertions = append(tempMap.Convertions,
				Convert{destination: nums[0], source: nums[1], rangeLength: nums[2]})
		} else {
			tempMap.name = line
		}
	}
	maps = append(maps, tempMap)
	return maps
}

func GetUInts32(text string) []uint32 {
	array := []uint32{}
	temp := uint32(0)
	i := 0
	for text[i] > '9' || text[i] < '0' {
		i++
	}
	for ; i < len(text); i++ {
		letter := text[i]
		if letter <= '9' && letter >= '0' {
			temp = temp*10 + uint32(letter) - uint32('0')
		} else {
			array = append(array, temp)
			temp = 0
		}
	}
	array = append(array, temp)
	return array
}

func min(array []uint32) uint32 {
	temp := array[0]
	for i := 1; i < len(array); i++ {
		if temp > array[i] {
			temp = array[i]
		}
	}
	return temp
}

func minPart2(array []uint32) uint32 {
	temp := array[0]
	for i := 0; i < len(array); i += 2 {
		if temp > array[i] {
			temp = array[i]
		}
	}
	return temp
}

func minNum(a uint32, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
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
