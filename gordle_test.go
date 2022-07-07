package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestCompareStrings(t *testing.T) {
	for _, tt := range []struct {
		name string
		input string
		reference string
		expected  []CharacterStatus
	}{
		{
			name: "Same_strigs",
			input: "TEST",
			reference: "TEST",
			expected:  []CharacterStatus{RIGHT, RIGHT, RIGHT, RIGHT},
		},
		{
			name: "Shuffeled",
			input: "STTE",
			reference: "TEST",
			expected:  []CharacterStatus{CONTAINS, CONTAINS, CONTAINS, CONTAINS},
		},
		{
			name: "Different",
			input: "OVAL",
			reference: "TEST",
			expected:  []CharacterStatus{WRONG, WRONG, WRONG, WRONG},
		},
		{
			name: "Complex",
			input: "TSAE",
			reference: "TEST",
			expected:  []CharacterStatus{RIGHT, CONTAINS, WRONG, CONTAINS},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res := CompareStrings(tt.input, tt.reference)

			if !reflect.DeepEqual(tt.expected, res) {
				t.Errorf("Comparition error: \n Expected %v \n Got %v", tt.expected, res)
			}
		})
	}
}

func TestBinSerach(t *testing.T) {
	file, err := os.Open("./dictionary/23.txt")

	if err != nil {
		t.Fatal("Could not open sample dictionary")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	arr := []string{}

	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	for pos, el := range arr {
		t.Run("Test/"+fmt.Sprint(pos), func(t *testing.T) {
			if !BinSearch(arr, el) {
				t.Errorf("Could not find %d string \"%s\"", pos, el)
			}
		})
	}
}
