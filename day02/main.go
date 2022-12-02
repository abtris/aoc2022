package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getStringRepresentation(shortcut string) string {
	switch shortcut {
	case "A", "X":
		return "Rock"
	case "B", "Y":
		return "Paper"
	case "C", "Z":
		return "Scissors"
	}
	return ""
}

func getScore(kind string) int {
	switch kind {
	case "Rock":
		return 1
	case "Paper":
		return 2
	case "Scissors":
		return 3
	}
	return 0
}

func getWinner(proposal, answer string) bool {
	if proposal == "Rock" && answer == "Scissors" {
		return true
	}
	if proposal == "Paper" && answer == "Rock" {
		return true
	}
	if proposal == "Scissors" && answer == "Paper" {
		return true
	}
	return false
}

func prepareData(filename string) map[int][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make(map[int][]string)
	index := 0
	for scanner.Scan() {
		s := scanner.Text()
		arr := strings.Split(s, " ")
		key := []string{getStringRepresentation(arr[0]), getStringRepresentation(arr[1])}
		data[index] = key
		index++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {
	data := prepareData("input")
	total := 0
	for _, v := range data {
		score := getScore(v[1])
		win := 0
		if v[0] == v[1] {
			win = 3
		}
		if getWinner(v[1], v[0]) {
			win = 6
		}
		total += score + win
	}
	fmt.Println(total)
}
