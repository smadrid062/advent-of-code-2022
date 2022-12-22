package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	path    = "./psr-guide-two"
	paper   = 1
	scissor = 2
	rock    = 3
	win     = 6
	lose    = 0
	draw    = 3
)

type stratergyGuide []string
type results []int

func findPlay(sg stratergyGuide) results {
	var rg results
	for _, val := range sg {
		var result int
		switch val {
		case "A X":
			result = lose + rock
		case "A Y":
			result = draw + paper
		case "A Z":
			result = win + scissor
		case "B X":
			result = lose + paper
		case "B Y":
			result = draw + scissor
		case "B Z":
			result = win + rock
		case "C X":
			result = lose + scissor
		case "C Y":
			result = draw + rock
		case "C Z":
			result = win + paper
		}
		rg = append(rg, result)
	}
	return rg
}

func getResult(sg stratergyGuide) results {
	var rg results
	for _, val := range sg {
		var result int
		switch val {
		case "A X":
			result = draw + paper
		case "A Y":
			result = win + scissor
		case "A Z":
			result = lose + rock
		case "B X":
			result = lose + paper
		case "B Y":
			result = draw + scissor
		case "B Z":
			result = win + rock
		case "C X":
			result = win + paper
		case "C Y":
			result = draw + scissor
		case "C Z":
			result = lose + rock
		}
		rg = append(rg, result)
	}
	return rg
}

func readFile(fileName string) (stratergyGuide, error) {
	fd, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	var rounds stratergyGuide
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {

		text := scanner.Text()
		if text == "" {
			continue
		}
		if len(text) != 3 {
			return nil, errors.New("error reading line " + text)
		}

		rounds = append(rounds, text)
	}
	return rounds, nil
}

func main() {
	stratGuide, err := readFile(path)
	if err != nil {
		return
	}
	rg := getResult(stratGuide)
	pg := findPlay(stratGuide)
	var rsum int
	var psum int
	for i := range stratGuide {
		rsum += rg[i]
		psum += pg[i]
	}
	fmt.Printf("assumed guide: %v\n", rsum)
	fmt.Printf("actual guide:  %v\n", psum)
}
