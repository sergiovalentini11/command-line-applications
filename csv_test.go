package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
	"testing/iotest"
)

func TestOperations(t *testing.T) {
	data := [][]float64{
		{10, 20, 15, 30, 45, 50, 100, 30},
		{5.5, 8, 2.2, 9.75, 8.45, 3, 2.5, 10.25, 4.75, 6.1, 7.67, 12.287, 5.47},
		{-10, -20},
		{102, 37, 44, 57, 67, 129},
	}

	// Testing scenarios for Operations
	testCases := []struct {
		name string
		op statsFunc
		exp []float64
	}{
		{"Sum", sum, []float64{300, 85.927, -30, 436}},
		{"Mean", mean, []float64{37.5, 6.609769230769231, -15, 72.666666666666666}},
	}

	// Operations test execution
	for _, tc := range testCases {
		for k, exp := range tc.exp {
			name := fmt.Sprintf("%sData%d", tc.name, k)
			t.Run(name, func(t *testing.T) {
				res := tc.op(data[k])

				if res != exp {
					t.Errorf("Expected %g, got %g instead", exp, res)
				}
			})
		}
	}
}

func TestCSV2Float(t *testing.T) {
	csvdata := `IP Address, Requests, Response Time
	192.168.0.199,2056,236
	192.168.0.88,899,220
	192.168.0.199,3054,226
	192.168.0.100,4133,218
	192.168.0.199,950,238
	`
}