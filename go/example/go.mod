module github.com/cschleiden/promptmd/example

go 1.23.4

replace github.com/cschleiden/promptmd => ../

require github.com/cschleiden/promptmd v0.0.0-20250222231703-f51a2d5dc8fa

require (
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
)

require (
	github.com/openai/openai-go v0.1.0-beta.2
	github.com/stretchr/testify v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
