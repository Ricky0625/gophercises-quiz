package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

const default_file = "problems.csv"

type QuizConfig struct {
	file string
}

// initialize quiz config
func (qc *QuizConfig) init() {
	flag.StringVar(&qc.file, "file", default_file, "path to file")
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

func main() {
	var config QuizConfig

	config.init()
	lines, err := config.readFile()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	fmt.Printf("Data: %+v\n", lines)
}
