package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const calorieList = "calorie-list"

func main() {
	elveList, err := fileParser(calorieList)
	if err != nil {
		log.Fatal(err)
	}
	maxCal := findHighest(elveList, 3)
	var sum int
	for _, max := range maxCal {
		fmt.Println(max)
		sum += max
	}
	fmt.Printf("sum: %v\n", sum)
}

type Elf struct {
	calories []int
	sum      int
}

func (e *Elf) getCalorieSum() {

	for _, val := range e.calories {
		e.sum += val
	}
}

func fileParser(fileName string) ([]Elf, error) {

	fd, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fd)

	var elveList []Elf
	for {
		var elf Elf
		var scanChecker bool
		for scanChecker = scanner.Scan(); scanChecker; scanChecker = scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				break
			}
			cal, err := strconv.Atoi(text)
			if err != nil {
				return nil, errors.New("fileparser: error parsing value")
			}
			elf.calories = append(elf.calories, cal)
		}
		if len(elf.calories) == 0 {
			if !scanChecker {
				break
			} else {
				continue
			}
		}
		elf.getCalorieSum()
		elveList = append(elveList, elf)
	}
	return elveList, nil

}

func findHighest(elves []Elf, nMax int) []int {
	max := make([]int, nMax)
	for _, elf := range elves {
		sum := elf.sum
		var prev int
		for i, val := range max {
			if prev > val {
				temp := max[i]
				max[i] = prev
				prev = temp
			}
			if sum > val {
				prev = max[i]
				max[i] = elf.sum
				sum = 0
			}
		}
	}
	return max
}
