package main

import (
	"flag"
	"fmt"
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

func main() {
	var config QuizConfig

	config.init()
	fmt.Println(config.file)
}
