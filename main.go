package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

const defaultFile = "problems.csv"

type QuizConfig struct {
	file string
}

type Problem struct {
	question string
	answer   string
}

// initialize quiz config
func (qc *QuizConfig) init() {
	flag.StringVar(&qc.file, "file", defaultFile, "path to file")
	flag.Parse()
}

func (qc *QuizConfig) readFile() ([][]string, error) {
	// open file
	f, err := os.Open(qc.file)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// read content from csv file
	csvReader := csv.NewReader(f)
	lines, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %w", err)
	}

	return lines, nil
}

/*
Do not use make([]problem, len(lines))
Issue: The original code used append to add elements to the problems slice, which caused a slice with twice the length of lines, leading to inefficient memory usage.
Solution: Use make with a length of 0 and a capacity of len(lines), or directly initialize the slice with the appropriate length to avoid unnecessary allocation.
*/

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, 0, len(lines))

	for _, line := range lines {
		problem := Problem{
			question: line[0],
			answer:   line[1],
		}
		problems = append(problems, problem)
	}
	return problems
}

func askQuestion(pb []Problem) (int, error) {
	var result int
	var input string

	for _, p := range pb {
		fmt.Printf("Q: %s\nA: ", p.question)
		_, err := fmt.Scan(&input)
		if err != nil {
			return 0, fmt.Errorf("failed to read input: %w", err)
		}

		if p.answer == input {
			result++
		}
	}

	return result, nil
}

func main() {
	var config QuizConfig

	config.init()
	lines, err := config.readFile()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	problems := parseLines(lines)
	result, err := askQuestion(problems)
	if err != nil {
		log.Fatalf("error running quiz: %v", err)
	}

	fmt.Printf("Result: %d/%d\n", result, len(lines))
}
