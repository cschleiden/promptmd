package promptmd

type PreparedPromptFunc func(vars Vars) ([]Message, error)

// Prepare compiles the prompt's messages for efficient template variable substitution.
//
// It processes each message in the prompt and returns a PreparedPromptFunc that can be
// used to generate the final set of messages with variables substituted. This approach
// allows for reusing the same prepared prompt with different variable values without
// re-parsing the templates.
//
// If the prompt contains no messages, it returns nil for both the function and error.
//
// The returned PreparedPromptFunc will:
// - Accept a Vars map containing variable values to substitute
// - Return the complete list of messages with variables replaced
// - Preserve the original role of each message
//
// Returns:
//   - PreparedPromptFunc: A function that accepts variables and returns messages
//   - error: Any error encountered during preparation
func (p *Prompt) Prepare() (PreparedPromptFunc, error) {
	if len(p.Messages) == 0 {
		return nil, nil
	}

	preparedMessages := make([]PreparedFunc, len(p.Messages))

	for i, message := range p.Messages {
		preparedMessages[i] = PrepareString(message.Message)
	}

	return func(vars Vars) ([]Message, error) {
		messages := make([]Message, len(preparedMessages))

		for i, prepared := range preparedMessages {
			msg, err := prepared(vars)
			if err != nil {
				return nil, err
			}

			messages[i] = Message{
				Role:    p.Messages[i].Role,
				Message: msg,
			}
		}

		return messages, nil
	}, nil
}
