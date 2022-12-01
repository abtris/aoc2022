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
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make(map[int]int)
	index := 1
	for scanner.Scan() {
		s := scanner.Text()
		if i, err := strconv.Atoi(s); err == nil {
			data[index] += i
		} else {
			index++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	max := 0
	simpleArray := []int{}
	for _, v := range data {
		simpleArray = append(simpleArray, v)
		if v > max {
			max = v
		}
	}
	sort.Ints(simpleArray)
	lastThree := len(simpleArray) - 3
	sum := 0
	for _, v := range simpleArray[lastThree:] {
		sum += v
	}
	fmt.Println(sum)
	fmt.Println(max)
}
