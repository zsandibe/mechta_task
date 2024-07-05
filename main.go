package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [json file path] [num of workers]")
		return
	}

	filePath := os.Args[1]
	numWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid number of workers:", err)
		return
	}

	if numWorkers < 1 || numWorkers > 5000 {
		fmt.Println("Number of workers must be between 1 and 5000")
		return
	}

	pairs, err := readAndUnmarshal(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	startTime := time.Now()
	totalSum := parallelSum(pairs, numWorkers)
	endTime := time.Now()
	fmt.Println("Total sum:", totalSum)
	fmt.Println("Time duration:", endTime.Sub(startTime))
}
