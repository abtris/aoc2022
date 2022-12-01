package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func getHungryElfCalories(data map[int]int) int {
	max := 0
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

func getHungryLastThreeElfCalories(data map[int]int) int {
	simpleArray := []int{}
	for _, v := range data {
		simpleArray = append(simpleArray, v)
	}
	sort.Ints(simpleArray)
	lastThree := len(simpleArray) - 3
	sum := 0
	for _, v := range simpleArray[lastThree:] {
		sum += v
	}
	return sum
}

func prepareData(filename string) map[int]int {
	file, err := os.Open(filename)
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
	return data
}

func main() {
	data := prepareData("input")
	fmt.Println(getHungryElfCalories(data))
	fmt.Println(getHungryLastThreeElfCalories(data))
}
