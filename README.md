# promptmd

## Go Module for Parsing `.prompt.md` Files

This repository includes a Go module for parsing LLM prompt files defined in `.prompt.md` files. The prompt files can contain an optional front matter defined in YAML. The markdown section of the prompt can contain just text, which means it's a prompt with a `user` role, or it can contain multiple prompts separated by `role:` on its own line. `role` can be `system`, `assistant`, or `user`.

### Usage

To use the Go module to parse `.prompt.md` files, follow the instructions below:

1. Create a new directory named `go` in your project.
2. Add the `parser.go` file to the `go` directory.
3. Import the `parser` package in your Go code.
4. Use the `ParsePromptFile` function to parse a `.prompt.md` file.

```go
package main

import (
	"fmt"
	"log"
	"path/to/your/project/go/parser"
)

func main() {
	promptFile, err := parser.ParsePromptFile("path/to/your/prompt.prompt.md")
	if err != nil {
		log.Fatalf("Error parsing prompt file: %v", err)
	}

	fmt.Printf("Front Matter: %v\n", promptFile.FrontMatter)
	for _, prompt := range promptFile.Prompts {
		fmt.Printf("Role: %s, Content: %s\n", prompt.Role, prompt.Content)
	}
}
```

### Example `.prompt.md` File

```markdown
---
title: Example Prompt
description: This is an example prompt file.
---

system:
You are a helpful assistant.

user:
What is the weather like today?

assistant:
The weather is sunny with a high of 75 degrees.
```
