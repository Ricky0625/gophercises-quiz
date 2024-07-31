package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	defaultFile      = "problems.csv"
	defaultTimeLimit = 30
	enableShuffle    = false
)

type QuizConfig struct {
	file      string
	timeLimit int
	shuffle   bool
}

type Problem struct {
	question string
	answer   string
}

// initialize quiz config
func (qc *QuizConfig) init() {
	flag.StringVar(&qc.file, "file", defaultFile, "path to file")
	flag.StringVar(&qc.file, "f", defaultFile, "path to file (shorthand)")
	flag.IntVar(&qc.timeLimit, "limit", defaultTimeLimit, "time limit for quiz")
	flag.IntVar(&qc.timeLimit, "l", defaultTimeLimit, "time limit for quiz (shorthand)")
	flag.BoolVar(&qc.shuffle, "shuffle", enableShuffle, "shuffle questions")
	flag.BoolVar(&qc.shuffle, "s", enableShuffle, "shuffle questions (shorthand)")
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

// read input using channel
// takes in a write-only channel, scan for input, send data through channel
func readInput(ansCh chan<- string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ansCh <- strings.TrimSpace(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error reading input: %v", err)
	}
	close(ansCh) // close in case of error
}

// ask questions and calculate score
func askQuestions(pb []Problem, timeLimit int) (int, error) {
	var score int
	ansCh := make(chan string)
	countdown := time.NewTimer(time.Duration(timeLimit) * time.Second)

	go readInput(ansCh)

	for _, p := range pb {
		fmt.Printf("Q: %s\nA: ", p.question)

		select {
		case <-countdown.C:
			fmt.Printf("\nTime's up!\n")
			return score, nil
		case answer := <-ansCh:
			if p.answer == answer {
				score++
			}
		}
	}

	return score, nil
}

func shuffleQuestions(pb []Problem) {
	rand.Shuffle(len(pb), func(i, j int) {
		pb[i], pb[j] = pb[j], pb[i]
	})
}

func main() {
	var config QuizConfig

	config.init()
	lines, err := config.readFile()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	problems := parseLines(lines)

	if config.shuffle {
		shuffleQuestions(problems)
	}

	score, err := askQuestions(problems, config.timeLimit)
	if err != nil {
		log.Fatalf("error running quiz: %v", err)
	}

	fmt.Printf("Score: %d/%d\n", score, len(problems))
}
