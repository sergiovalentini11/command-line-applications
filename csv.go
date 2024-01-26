package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)


// Sums all the numbers in a given slice
func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}
	return sum
}


// Uses sum to calculate the mean of a given slice
func mean(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// use this type for generic statistical functions
type statsFunc func(data []float64) float64

