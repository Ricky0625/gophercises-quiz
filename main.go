package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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
	flag.StringVar(&qc.file, "f", defaultFile, "path to file (shorthand)")
	flag.Parse()
}

// open file and load csv file content
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

// parse csv lines into problems
func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, 0, len(lines))

	for _, line := range lines {
		if len(line) < 2 {
			continue // skip malformed lines
		}

		problem := Problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
		problems = append(problems, problem)
	}
	return problems
}

// ask questions and calculate score
func askQuestion(pb []Problem) (int, error) {
	var score int
	scanner := bufio.NewScanner(os.Stdin)

	for _, p := range pb {
		fmt.Printf("Q: %s\nA: ", p.question)
		if !scanner.Scan() {
			return 0, fmt.Errorf("failed to read input: %w", scanner.Err())
		}
		input := strings.TrimSpace(scanner.Text())

		if p.answer == input {
			score++
		}
	}

	return score, nil
}

func main() {
	var config QuizConfig

	config.init()
	lines, err := config.readFile()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	problems := parseLines(lines)
	score, err := askQuestion(problems)
	if err != nil {
		log.Fatalf("error running quiz: %v", err)
	}

	fmt.Printf("Score: %d/%d\n", score, len(problems))
}
