package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
)

func parseFlags() int {
	var numGoroutines int
	flag.IntVar(&numGoroutines, "goroutines", runtime.NumCPU(), "number of goroutines to use")
	flag.Parse()
	return numGoroutines
}

func readAndUnmarshalJson(fileName string) []Data {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}
	var dataArray []Data
	if err = json.Unmarshal(content, &dataArray); err != nil {
		fmt.Printf("Error unmarshaling json: %s\n", err)
		os.Exit(1)
	}
	return dataArray
}

func processItemsConcurrently(dataArray []Data, numGoroutines int) chan int {
	jobs := make(chan Data, len(dataArray))
	results := make(chan int, len(dataArray))

	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go worker(jobs, results, &wg)
	}

	for _, data := range dataArray {
		jobs <- data
	}
	close(jobs)

	wg.Wait()
	close(results)
	return results
}

func worker(jobs <-chan Data, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job.A + job.B
	}
}

func aggregateResults(results chan int) int {
	totalSum := 0
	for result := range results {
		totalSum += result
	}
	return totalSum
}
