/**
*	go test -v -run TestAdd
 */

package main

import (
	"io/ioutil"
	"testing"
)

type AddResult struct {
	x        int
	y        int
	expected int
}

var AddResults = []AddResult{
	{1, 1, 2},
}

func TestAdd(t *testing.T) {

	for _, test := range AddResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Result is not as expected!")
		}
	}
}

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}
	if string(data) != "hello world" {
		t.Fatal("Contents does not match expected.")
	}
}
