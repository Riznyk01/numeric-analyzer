package main

import (
	"bufio"
	"flag"
	"fmt"
	"numeric-analyzer/internal/analyzer"
	"os"
	"strconv"
	"time"
)

func main() {
	pathPtr := flag.String("path", "", "the path to the file")
	flag.Parse()
	t := time.Now()
	data, err := readFile(*pathPtr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	minimum, maximum, median, avg, err := analyzer.ProcessNumericData(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Printf("max: %d\n", maximum)
	fmt.Printf("min: %d\n", minimum)
	fmt.Printf("median: %0.2f\n", median)
	fmt.Printf("avg: %0.2f\n", avg)

	increasingSequence, err := analyzer.FindSequences(data, true)
	printSequence("increasing", increasingSequence, err)
	decreasingSequence, err := analyzer.FindSequences(data, false)
	printSequence("decreasing", decreasingSequence, err)

	fmt.Printf("time of executing: %v\n", time.Since(t))
}

func printSequence(sequenceType string, sequence [][]int, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if sequence != nil {
		fmt.Printf("max %s sequence: %v\n", sequenceType, sequence)
	} else {
		fmt.Printf("max %s sequence doesn't exist in this file\n", sequenceType)
	}
}

func readFile(path string) ([]int, error) {
	dataArr := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open the file %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("error while converting numbers from the file: %v", err)
		}
		dataArr = append(dataArr, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error during file scanning: %v", err)
	}

	if len(dataArr) == 0 {
		return nil, fmt.Errorf("the file is empty")
	}

	return dataArr, nil
}
