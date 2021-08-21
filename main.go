package parser

import (
	"github.com/alecthomas/participle/v2"
)

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
	err = parser.ParseString("", rawCode, rootNode)
	if err != nil {
		return nil, err
	}

	err = postProcess(rootNode)
	if err != nil {
		return nil, err
	}

	return rootNode, nil
}
