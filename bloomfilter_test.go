package bloomfilter

import (
	"fmt"
	"testing"
)

func TestFindWords(t *testing.T) {
	testCases := []struct{
		expected bool
		phraseGiven string
	}{
		{true,  "Peter"},
		{true,  "Peter is crazy"},
		{false, "Peter is creisi"},
		{true,  "chair"},
		{false, "clavo"},
		{true,  "nail"},
		{true,  "child"},
		{false, "chavo"},
		{false,  "koson is an artist"},
		{true,  "Peter is an artist"},
	}
	LoadMap()
	for i, testCase := range testCases {
		testName := fmt.Sprintf("Test-FindWords-%d", i+1)
		t.Run(testName, func(t *testing.T) {
			actual := FindWords(testCase.phraseGiven)
			if actual != testCase.expected {
				t.Errorf("Failed: expected %v, actual %v", testCase.expected, actual)
			}
		})
	}
}
