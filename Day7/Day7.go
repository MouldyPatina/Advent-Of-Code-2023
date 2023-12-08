package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	hands := readFile("Real")
	fmt.Println("This is the solution to Part 1:")
	fmt.Println(Part1Solver(hands))
	fmt.Println("This is the solution to Part 2:")
	fmt.Println(Part2Solver(hands))
}

func Part1Solver(hands []string) int {
	allHands := make([]Hand, len(hands));

	for index, hand := range hands {
		tempHand := Hand{};
		tempHand.card = hand[:5];
		tempHand.bid = getInt(hand[6:])
		tempHand.strength = getStrength(tempHand.card);
		allHands[index] = tempHand;
	}
	sort.Sort(ByHand(allHands));
	output := 0;
	for index, card := range allHands {
		output += card.bid * (index + 1); 
	}

	return output
}

func Part2Solver(hands []string) int {
	allHands := make([]Hand, len(hands));

	for index, hand := range hands {
		tempHand := Hand{};
		tempHand.card = hand[:5];
		tempHand.bid = getInt(hand[6:])
		tempHand.strength = getStrengthJ(tempHand.card);
		allHands[index] = tempHand;
	}
	sort.Sort(ByJoker(allHands));
	output := 0;
	for index, card := range allHands {
		output += card.bid * (index + 1); 
	}

	return output
}

type Hand struct {
	card 		string
	bid		int
	strength	int
}

type ByHand []Hand

func (a ByHand) Len() int {return len(a)}
func (a ByHand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByHand) Less(i, j int) bool {
	if a[i].strength < a[j].strength {
		return true;
	} else if a[i].strength == a[j].strength {
		for w := 0; w < 5; w++ {
			if CardStrength(a[i].card[w]) < CardStrength(a[j].card[w]) {
				return true;
			} else if CardStrength(a[j].card[w]) < CardStrength(a[i].card[w]) {
				return false;
			}
		}
	}
	return false;
}

type ByJoker []Hand

func (a ByJoker) Len() int {return len(a)}
func (a ByJoker) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByJoker) Less(i, j int) bool {
	if a[i].strength < a[j].strength {
		return true;
	} else if a[i].strength == a[j].strength {
		for w := 0; w < 5; w++ {
			if CardStrengthJ(a[i].card[w]) < CardStrengthJ(a[j].card[w]) {
				return true;
			} else if CardStrengthJ(a[j].card[w]) < CardStrengthJ(a[i].card[w]) {
				return false;
			}
		}
	}
	return false;
}

func CardStrengthJ(card byte) int {
	order := "J23456789TQKA";
	for i := 0; i < len(order); i++ {
		if order[i] == card {
			return i;
		}
	}
	log.Println("Invalid card", string(card));
	return -1;
}

func CardStrength(card byte) int {
	order := "23456789TJQKA";
	for i := 0; i < len(order); i++ {
		if order[i] == card {
			return i;
		}
	}
	log.Println("Invalid card", string(card));
	return -1;
}

func getInt(str string) int {
	output := 0;
	for i := 0; i < len(str); i++ {
		output = output * 10 + int(str[i]) - int('0')
	}
	return output;
}

func getStrengthJ(cards string) int {
	jokers := 0; 
	used := "";
	count := []int{0};
	for i := 0; i < len(cards); i++ {
		if cards[i] == 'J' {
			jokers++;
		} else if !Contains(used, cards[i]) {
			count = append( count, Count(cards, cards[i]));
			used = cards[:i + 1];
		}
	}
	sort.Slice(count, func(i, j int) bool { return count[i] > count[j] });
	switch jokers + count[0] {
	case 5:
		return 6;
	case 4:
		return 5;
	case 3:
		if count[1] == 2{
			return 4;
		}
		return 3;
	case 2:
		if count[1] == 2 {
			return 2;
		}
		return 1;
	case 1:
		return 0;
	default:
		log.Println("invalid hand");
		return -1;
	}
}

func getStrength(cards string) int {
	used := "";
	count := [5]int{};
	for i := 0; i < len(cards); i++ {
		if !Contains(used, cards[i]) {
			count[Count(cards, cards[i]) - 1]++;
			used = cards[:i + 1];
		}
	}
	if count [4] == 1 {
		return 6;
	} else if count[3] == 1 {
		return 5;
	} else if count[2] == 1 {
		if count[1] == 1 {
			return 4;
		} else {
			return 3;
		}
	} else if count[1] == 2 {
		return 2;
	} else if count[1] == 1 {
		return 1;
	} else {
		return 0;
	}
}

func Contains(str string, char byte) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			return true;
		}
	}
	return false;
}

func Count(str string, char byte) int {
	count := 0;
	for i := 0; i < len(str); i++ {
		if str[i] == char {
			count++;
		}
	}
	return count;
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
