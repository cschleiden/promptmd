package main

import (
	"context"
	"log"
	"os"

	"github.com/cschleiden/promptmd"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	// Read prompt
	content, err := os.ReadFile("recipe.prompt.md")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse
	p, err := promptmd.Parse(string(content))
	if err != nil {
		log.Fatalf("Error parsing prompt: %v", err)
	}

	// Prepare the prompt for easier variable substitution, this needs to be done only once
	pt, err := p.Prepare()
	if err != nil {
		log.Fatalf("Error preparing prompt: %v", err)
	}

	// Replace any placeholders in the parsed prompt. Beware of any untrusted input here.
	msg, err := pt(promptmd.Vars{
		"ingredient": "broccoli",
	})
	if err != nil {
		log.Fatalf("Error processing prompt: %v", err)
	}

	// Prompt GitHub Models with the prepared prompt
	oaimsg := make([]openai.ChatCompletionMessageParamUnion, len(msg))
	for i, m := range msg {
		switch m.Role {
		case promptmd.RoleUser:
			oaimsg[i] = openai.UserMessage(m.Message)
		case promptmd.RoleSystem:
			oaimsg[i] = openai.SystemMessage(m.Message)
		}
	}

	client := openai.NewClient(
		option.WithBaseURL("https://models.inference.ai.azure.com"),
		option.WithAPIKey(os.Getenv("GITHUB_TOKEN")),
	)
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: oaimsg,
		Model:    p.Metadata["model"].(string),
	})
	if err != nil {
		panic(err.Error())
	}

	println(chatCompletion.Choices[0].Message.Content)
}
