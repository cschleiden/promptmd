package promptmd

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
)

var Roles = []Role{RoleUser, RoleAssistant, RoleSystem}

type Message struct {
	Role    Role   `json:"role"`
	Message string `json:"message"`
}

type Prompt struct {
	Metadata map[string]any `json:"metadata"`

	Messages []Message `json:"messages"`
}
