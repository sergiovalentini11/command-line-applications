package main

import (
	"bytes"
	"errors"
	"os"
	"testing"
)

func TestRun(t *testing.T) {

	// Test cases for run
	testCases := []struct {
		name string
		col int
		op string
		exp string
		files []string
		expErr error
	}{
		{name: "RunMean1File", col: 3, op: "mean", exp: "227.6\n",
			files: []string{"./testdata/example.csv"},
			expErr: nil,
		},
		{name: "RunMeanMultiFiles", col: 3, op: "mean", exp: "233.84\n",
			files: []string{"./testdata/example.csv", "./testdata/example2.csv"},
			expErr: nil,
		},
		{name: "RunFailRead", col: 2, op: "mean", exp: "",
			files: []string{"./testdata/example.csv", "./testdata/fakefile.csv"},
			expErr: os.ErrNotExist,
		},
		{name: "RunFailColumn", col: 0, op: "mean", exp: "",
			files: []string{"./testdata/example.csv"},
			expErr: ErrInvalidColumn,
		},
		{name: "RunFailNoFiles", col: 2, op: "mean", exp: "",
			files: []string{},
			expErr: ErrNoFiles,
		},
		{name: "RunFailOperation", col: 2, op: "invalid", exp:"",
			files: []string{"./testdata/example.csv"},
			expErr: ErrInvalidOperation,
		},
	}

	// Run test executing
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var res bytes.Buffer
			err := run(tc.files, tc.op, tc.col, &res)

			if tc.expErr != nil {
				if err == nil {
					t.Errorf("Expected error. Got nil instead")
				}
				if ! errors.Is(err, tc.expErr) {
					t.Errorf("Expected error %q, got %q instead", tc.expErr, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected rror: %q", err)
			}
			if res.String() != tc.exp {
				t.Errorf("Expected %q, got %q instead", tc.exp, &res)
			}
		})
	}
}