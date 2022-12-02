package main

import (
	"bufio"
	"errors"
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

func getScore(kind string) (int, error) {
	switch kind {
	case "Rock":
		return 1, nil
	case "Paper":
		return 2, nil
	case "Scissors":
		return 3, nil
	}
	return 0, errors.New("Nothing selected")
}

func getWinnerScore(proposal, answer string) (int, error) {
	if proposal == "Rock" {
		switch answer {
		case "Rock":
			return 3, nil
		case "Paper":
			return 6, nil
		case "Scissors":
			return 0, nil
		}
	}
	if proposal == "Paper" {
		switch answer {
		case "Rock":
			return 0, nil
		case "Paper":
			return 3, nil
		case "Scissors":
			return 6, nil
		}
	}
	if proposal == "Scissors" {
		switch answer {
		case "Rock":
			return 6, nil
		case "Paper":
			return 0, nil
		case "Scissors":
			return 3, nil
		}
	}
	return 0, errors.New("Wrong combination!")
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
		score, err := getScore(v[0])
		if err != nil {
			log.Fatalln(err)
		}
		win, err := getWinnerScore(v[0], v[1])
		if err != nil {
			log.Fatalln(err)
		}
		total += score + win
	}
	fmt.Println(total)
}
