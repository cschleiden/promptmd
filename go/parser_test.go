package promptmd_test

import (
	"os"
	"testing"

	"github.com/cschleiden/promptmd"
)

func TestParsePromptFile_EmptyFile(t *testing.T) {
	filePath := createTempFile(t, "")
	defer os.Remove(filePath)

	promptFile, err := promptmd.ParsePromptFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing prompt file: %v", err)
	}

	if len(promptFile.FrontMatter) != 0 {
		t.Errorf("Expected empty front matter, got: %v", promptFile.FrontMatter)
	}

	if len(promptFile.Prompts) != 0 {
		t.Errorf("Expected no prompts, got: %v", promptFile.Prompts)
	}
}

func TestParsePromptFile_OnlyFrontMatter(t *testing.T) {
	content := `---
title: Example Prompt
description: This is an example prompt file.
---`
	filePath := createTempFile(t, content)
	defer os.Remove(filePath)

	promptFile, err := promptmd.ParsePromptFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing prompt file: %v", err)
	}

	if len(promptFile.FrontMatter) == 0 {
		t.Errorf("Expected front matter, got: %v", promptFile.FrontMatter)
	}

	if len(promptFile.Prompts) != 0 {
		t.Errorf("Expected no prompts, got: %v", promptFile.Prompts)
	}
}

func TestParsePromptFile_OnlyPrompts(t *testing.T) {
	content := `system:
You are a helpful assistant.

user:
What is the weather like today?

assistant:
The weather is sunny with a high of 75 degrees.`
	filePath := createTempFile(t, content)
	defer os.Remove(filePath)

	promptFile, err := promptmd.ParsePromptFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing prompt file: %v", err)
	}

	if len(promptFile.FrontMatter) != 0 {
		t.Errorf("Expected no front matter, got: %v", promptFile.FrontMatter)
	}

	if len(promptFile.Prompts) != 3 {
		t.Errorf("Expected 3 prompts, got: %v", promptFile.Prompts)
	}
}

func TestParsePromptFile_FrontMatterAndPrompts(t *testing.T) {
	content := `---
title: Example Prompt
description: This is an example prompt file.
---

system:
You are a helpful assistant.

user:
What is the weather like today?

assistant:
The weather is sunny with a high of 75 degrees.`
	filePath := createTempFile(t, content)
	defer os.Remove(filePath)

	promptFile, err := promptmd.ParsePromptFile(filePath)
	if err != nil {
		t.Fatalf("Error parsing prompt file: %v", err)
	}

	if len(promptFile.FrontMatter) == 0 {
		t.Errorf("Expected front matter, got: %v", promptFile.FrontMatter)
	}

	if len(promptFile.Prompts) != 3 {
		t.Errorf("Expected 3 prompts, got: %v", promptFile.Prompts)
	}
}

func createTempFile(t *testing.T, content string) string {
	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatalf("Error writing to temp file: %v", err)
	}

	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Error closing temp file: %v", err)
	}

	return tmpFile.Name()
}
