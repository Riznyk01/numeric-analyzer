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
	} else {
		fmt.Printf("min number is: %d\n", analyzer.Min(data))
		fmt.Printf("max number is: %d\n", analyzer.Max(data))
		fmt.Printf("avg number is: %0.1f\n", analyzer.Avg(data))
		fmt.Printf("median number is: %d\n", analyzer.Median(data))
	}
	fmt.Printf("time of executing: %v\n", time.Since(t))
}
func readFile(path string) ([]int, error) {
	dataArr := make([]int, 0)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("error while reading numbers from the file %v", err)
		}
		dataArr = append(dataArr, num)
	}
	return dataArr, nil
}
