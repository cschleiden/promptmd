package promptmd

import (
	"encoding/json"
	"os"
	"testing"
)

type templateTest struct {
	Name   string         `json:"name"`
	Input  string         `json:"input"`
	Vars   map[string]any `json:"vars"`
	Output string         `json:"output"`
}

func Test_Template_Prepare(t *testing.T) {
	jsonFile, err := os.Open("../testdata/template/template.json")
	if err != nil {
		t.Fatalf("Failed to open testdata: %s", err)
	}
	defer jsonFile.Close()

	var tests []templateTest
	err = json.NewDecoder(jsonFile).Decode(&tests)
	if err != nil {
		t.Fatalf("Failed to decode testdata: %s", err)
	}

	// Run test cases
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			prepared := Prepare(test.Input)
			result, err := prepared(test.Vars)

			if err != nil {
				t.Fatalf("Prepare returned an error: %s", err)
			}

			if result != test.Output {
				t.Errorf("Expected %q, but got %q", test.Output, result)
			}
		})
	}
}
