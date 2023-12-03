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
	calibrations := readFile("Test");
	for _, line := range calibrations {
		fmt.Println(line)
	}
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
