package main

import (
	"fmt"
	"time"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	numGoroutines := parseFlags()
	dataArray := readAndUnmarshalJson("data.json")
	results := processItemsConcurrently(dataArray, numGoroutines)
	startTime := time.Now()
	totalSum := aggregateResults(results)
	endTime := time.Now()
	fmt.Printf("Total sum: %d\n", totalSum)
	fmt.Printf("Time duration: %v\n", endTime.Sub(startTime))
}
