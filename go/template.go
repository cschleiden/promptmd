package promptmd

import (
	"regexp"
	"strconv"
)

type Vars map[string]any

type PreparedFunc func(vars Vars) (string, error)

func Prepare(template string) PreparedFunc {
	segments := parseTemplate(template)

	return func(vars Vars) (string, error) {
		result := ""
		for _, segment := range segments {
			if !segment.IsVariable {
				result += segment.Content
				continue
			}

			// If there is no value a variable skip it
			if val, ok := vars[segment.Content]; ok {
				switch v := val.(type) {
				case string:
					result += v
				case nil:
					result += ""
				case bool:
					if v {
						result += "true"
					} else {
						result += "false"
					}
				case int:
					result += strconv.FormatInt(int64(v), 10)
				case int64:
					result += strconv.FormatInt(v, 10)
				case float64:
					result += strconv.FormatFloat(v, 'f', -1, 64)
				default:
					if stringer, ok := val.(interface{ String() string }); ok {
						result += stringer.String()
					}
				}
			}
		}

		return result, nil
	}
}

type templateSegment struct {
	Content    string
	IsVariable bool
}

var variableRegex = regexp.MustCompile(`{{(.*?)}}`)

func parseTemplate(template string) []templateSegment {
	segments := make([]templateSegment, 0)

	matches := variableRegex.FindAllStringSubmatchIndex(template, -1)

	current := 0
	for _, match := range matches {
		variableStart := match[0]
		variableEnd := match[1]
		contentStart := match[2]
		contentEnd := match[3]

		if current < variableStart {
			segments = append(segments, templateSegment{
				Content:    template[current:variableStart],
				IsVariable: false,
			})
		}

		segments = append(segments, templateSegment{
			Content:    template[contentStart:contentEnd],
			IsVariable: true,
		})

		current = variableEnd
	}

	if current < len(template) {
		segments = append(segments, templateSegment{
			Content:    template[current:],
			IsVariable: false,
		})
	}

	return segments
}
