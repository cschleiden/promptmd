package promptmd

import (
	"bufio"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Prompt struct {
	Role    string
	Content string
}

type PromptFile struct {
	FrontMatter map[string]interface{}
	Prompts     []Prompt
}

func ParsePromptFile(filePath string) (*PromptFile, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var frontMatterLines []string
	var prompts []Prompt
	var currentRole string
	var currentContent strings.Builder
	inFrontMatter := false
	inPrompt := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "---" {
			if !inFrontMatter {
				inFrontMatter = true
			} else {
				inFrontMatter = false
				inPrompt = true
			}
			continue
		} else {
			if !inFrontMatter && !inPrompt {
				inPrompt = true
			}
		}

		if inFrontMatter {
			frontMatterLines = append(frontMatterLines, line)
		} else if inPrompt {
			if strings.HasSuffix(line, ":") && (strings.HasPrefix(line, "user") || strings.HasPrefix(line, "system") || strings.HasPrefix(line, "assistant")) {
				if currentRole != "" {
					prompts = append(prompts, Prompt{
						Role:    currentRole,
						Content: currentContent.String(),
					})
					currentContent.Reset()
				}
				currentRole = strings.TrimSuffix(line, ":")
			} else {
				currentContent.WriteString(line + "\n")
			}
		}
	}

	if currentRole != "" {
		prompts = append(prompts, Prompt{
			Role:    currentRole,
			Content: currentContent.String(),
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	frontMatter := make(map[string]interface{})
	if len(frontMatterLines) > 0 {
		err = yaml.Unmarshal([]byte(strings.Join(frontMatterLines, "\n")), &frontMatter)
		if err != nil {
			return nil, err
		}
	}

	return &PromptFile{
		FrontMatter: frontMatter,
		Prompts:     prompts,
	}, nil
}
