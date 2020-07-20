package main

import (
	"fmt"
	"testing"
)

func TestReadLines(t *testing.T) {
	var tests = []struct {
		fname string
		idx int
		want string
	}{
		{"sample/test-read.txt", 0, "This is a test read"},
		{"sample/test-read.txt", 1, "This is some content"},
		{"sample/test-read.txt", 2, "This is some other stuff"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.idx)
		t.Run(testname, func(t *testing.T) {
			ans := ReadLines(tt.fname)[tt.idx]
			if ans != tt.want {
				t.Errorf("got `%s`, want `%s`", ans, tt.want)
			}
		})
	}
}
