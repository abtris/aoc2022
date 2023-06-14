package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/juliangruber/go-intersect/v2"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	part2Count := 0
	for scanner.Scan() {
		s := scanner.Text()
		members := strings.Split(s, ",")
		a := getArrayFromString(members[0])
		b := getArrayFromString(members[1])
		res := intersect.Simple(a, b)
		if len(res) != 0 {
			part2Count++
			if len(res) == len(a) || len(res) == len(b) {
				count++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	fmt.Println(part2Count)
}

func getArrayFromString(s string) []int {
	borders := strings.Split(s, "-")
	var output []int
	lowerBound, _ := strconv.Atoi(borders[0])
	highBound, _ := strconv.Atoi(borders[1])

	for i := lowerBound; i <= highBound; i++ {
		output = append(output, i)
	}
	return output
}
