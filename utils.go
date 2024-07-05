package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

func readAndUnmarshal(filePath string) ([]Data, error) {
	if filepath.Ext(filePath) != ".json" {
		return nil, errors.New("File is not a JSON file")
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %w", err)
	}

	var pairs []Data
	if err := json.Unmarshal(data, &pairs); err != nil {
		return nil, fmt.Errorf("Error unmarshalling json: %w", err)
	}

	return pairs, nil
}

func parallelSum(pairs []Data, goroutineCount int) int {
	results := make(chan int, goroutineCount)
	var wg sync.WaitGroup

	chunkSize := (len(pairs) + goroutineCount - 1) / goroutineCount
	for i := 0; i < len(pairs); i += chunkSize {
		end := i + chunkSize
		if end > len(pairs) {
			end = len(pairs)
		}

		wg.Add(1)
		go func(pairsChunk []Data) {
			defer wg.Done()
			sum := 0
			for _, pair := range pairsChunk {
				sum += pair.A + pair.B
			}
			results <- sum
		}(pairs[i:end])
	}

	wg.Wait()
	close(results)

	totalSum := 0
	for sum := range results {
		totalSum += sum
	}

	return totalSum
}
