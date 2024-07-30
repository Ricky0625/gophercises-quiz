package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

const defaultFile = "problems.csv"

type quizConfig struct {
	file string
}

type problem struct {
}

// initialize quiz config
func (qc *quizConfig) init() {
	flag.StringVar(&qc.file, "file", defaultFile, "path to file")
	flag.Parse()
}

func (qc *quizConfig) readFile() ([][]string, error) {
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
	var config quizConfig

	config.init()
	lines, err := config.readFile()
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	fmt.Printf("Data: %+v\n", lines)
}
