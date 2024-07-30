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
Do not use make([]problem, 0, len(lines))
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

func main() {
	var config QuizConfig

	config.init()
	lines, err := config.readFile()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	problems := parseLines(lines)
	for _, problem := range problems {
		fmt.Printf("q: %s, a: %s\n", problem.question, problem.answer)
	}
}
