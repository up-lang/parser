package parser

import (
	"github.com/alecthomas/participle/v2"
)

func ParseFromFile(filePath string) (Up, error) {
	code, err := readFile(filePath)
	if err != nil {
		return Up{}, err
	}
	return Parse(code)
}

func Parse(rawCode string) (Up, error) {
	parser, err := participle.Build(&Up{})
	if err != nil {
		return Up{}, err
	}

	rootNode := Up{}
	err = parser.ParseString("", rawCode, rootNode)
	if err != nil {
		return Up{}, err
	}

	return rootNode, nil
}
