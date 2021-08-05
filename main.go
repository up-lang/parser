package parser

import (
	"github.com/alecthomas/participle/v2"
)

func ParseFromFile(filePath string) (UpRoot, error) {
	code, err := readFile(filePath)
	if err != nil {
		return UpRoot{}, err
	}
	return Parse(code)
}

func Parse(rawCode string) (UpRoot, error) {
	padded := pad(rawCode)

	collapsed := collapseWhitespace(padded)

	parser, err := participle.Build(&UpRoot{})
	if err != nil {
		return UpRoot{}, err
	}

	rootNode := UpRoot{}
	err = parser.ParseString("", collapsed, rootNode)
	if err != nil {
		return UpRoot{}, err
	}

	return rootNode, nil
}

func pad(raw string) string {
	return PadRegex.ReplaceAllString(raw, " $& ")
}
