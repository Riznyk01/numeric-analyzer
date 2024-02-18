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
	minimum, maximum, median, avg := analyzer.ProcessNumericData(data)

	fmt.Printf("min number is: %d\n", minimum)
	fmt.Printf("min number is: %d\n", maximum)
	fmt.Printf("median number is: %0.1f\n", median)
	fmt.Printf("avg number is: %0.1f\n", avg)

	increasingSequence := analyzer.FindSequences(data, true)
	if increasingSequence != nil {
		fmt.Printf("max increasing sequence is: %v\n", increasingSequence)
	} else {
		fmt.Println("max increasing sequence doesn't exist in this file")
	}

	decreasingSequence := analyzer.FindSequences(data, false)
	if decreasingSequence != nil {
		fmt.Printf("max decreasing sequence is: %v\n", decreasingSequence)
	} else {
		fmt.Println("max decreasing sequence doesn't exist in this file")
	}

	fmt.Printf("time of executing: %v\n", time.Since(t))
}
func readFile(path string) ([]int, error) {
	dataArr := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error while opening the file %v", err)
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
