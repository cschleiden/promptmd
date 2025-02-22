package promptmd

import (
	"regexp"
	"strings"

	"github.com/stretchr/testify/assert/yaml"
)

func Parse(content string) (*Prompt, error) {
	p := &Prompt{}

	// Parse metadata
	// Check if there is frontmatter
	if strings.HasPrefix(content, "---") {
		end := strings.Index(content[3:], "---")
		if end > 0 {
			// Parse frontmatter
			frontmatter := content[3 : end+3]
			var metadata map[string]any
			err := yaml.Unmarshal([]byte(frontmatter), &metadata)
			if err != nil {
				return nil, err
			}

			p.Metadata = metadata

			// Remove frontmatter from content
			content = content[end+6:]
		}
	}

	// Parse content
	messages, err := parseMessages(content)
	if err != nil {
		return nil, err
	}
	p.Messages = messages

	return p, nil
}

var roleRegex = regexp.MustCompile(
	`(?m)^\s*#?\s*(` + strings.Join(stringSlice(Roles), "|") + `)\s*:\s*$`,
)

func stringSlice(roles []Role) []string {
	ss := make([]string, len(roles))
	for i, r := range roles {
		ss[i] = string(r)
	}
	return ss
}

func parseMessages(content string) ([]Message, error) {
	// Parse the template(s)
	messages := []Message{}

	matches := roleRegex.FindAllStringSubmatchIndex(content, -1)

	// If there are no roles, assume the entire content is a system prompt
	if len(matches) == 0 {
		messages = append(messages, Message{
			Role:    RoleSystem,
			Message: strings.TrimSpace(content),
		})
		return messages, nil
	}

	// Iterate over the matches and extract the content
	for i, match := range matches {
		// Extract the role
		role := Role(strings.TrimSpace(content[match[2]:match[3]]))

		var contentStart int
		if i == len(matches)-1 {
			contentStart = match[1]
		} else {
			contentStart = match[1]
		}

		var contentEnd int
		if i == len(matches)-1 {
			contentEnd = len(content)
		} else {
			contentEnd = matches[i+1][0]
		}

		messageContent := content[contentStart:contentEnd]

		// Extract the content
		messages = append(messages, Message{
			Role:    role,
			Message: strings.TrimSpace(messageContent),
		})
	}

	return messages, nil
}
