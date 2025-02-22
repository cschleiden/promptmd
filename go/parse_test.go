package promptmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	testDataDir := filepath.Join("..", "testdata", "parse") // Assuming testdata is one level up
	testFiles, err := os.ReadDir(testDataDir)
	if err != nil {
		t.Fatalf("Failed to read testdata directory: %v", err)
	}

	var promptFiles []map[string]string

	for _, file := range testFiles {
		if strings.HasSuffix(file.Name(), ".prompt.md") {
			name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			promptFiles = append(promptFiles, map[string]string{
				"name": name,
				"file": file.Name(),
			})
		}
	}

	for _, testFile := range promptFiles {
		name := testFile["name"]
		file := testFile["file"]

		t.Run(name, func(t *testing.T) {
			promptFile := filepath.Join(testDataDir, file)
			expectedFile := filepath.Join(testDataDir, strings.Replace(file, ".prompt.md", ".expected.json", 1))

			promptContentBytes, err := os.ReadFile(promptFile)
			if err != nil {
				t.Fatalf("Failed to read prompt file: %v", err)
			}
			promptContent := string(promptContentBytes)

			expectedContentBytes, err := os.ReadFile(expectedFile)
			if err != nil {
				t.Fatalf("Failed to read expected file: %v", err)
			}
			expectedContent := string(expectedContentBytes)

			var expected *Prompt
			err = json.Unmarshal([]byte(expectedContent), &expected)
			if err != nil {
				t.Fatalf("Failed to unmarshal expected JSON: %v", err)
			}

			result, err := Parse(promptContent)

			require.NoError(t, err)
			require.Equal(t, expected, result)
		})
	}
}
