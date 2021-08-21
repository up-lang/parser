package parser

import (
	"github.com/alecthomas/participle/v2"
	"regexp"
	"strings"
)

//goland:noinspection GoUnusedExportedFunction
func ParseFromFile(filePath string) (*Up, error) {
	code, err := readFile(filePath)
	if err != nil {
		return nil, err
	}
	return Parse(code)
}

func Parse(rawCode string) (*Up, error) {
	parser, err := participle.Build(&Up{})
	if err != nil {
		return nil, err
	}

	rootNode := &Up{}
	err = parser.ParseString("", RemoveComments(rawCode), rootNode)
	if err != nil {
		return nil, err
	}

	err = postProcess(rootNode)
	if err != nil {
		return nil, err
	}

	return rootNode, nil
}

func RemoveComments(raw string) string {
	// block comments
	// [\s\S] is like . except it also matches newlines
	// replace with nothing to entirely remove
	noBlockComments := regexp.MustCompile("~~~[\\s\\S]*~~~").ReplaceAllString(raw, "")

	// line comments
	lines := strings.Split(noBlockComments, "\n")
	sb := strings.Builder{}
	for i := range lines {

		for i2 := range lines[i] {

			if lines[i][i2] == '~' {
				break
			}

			sb.WriteRune(rune(lines[i][i2]))
		}

		if i+1 < len(lines) { // if not the last item
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}
