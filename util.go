package parser

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	LineEndingRegex = regexp.MustCompile("(?:\r?\n)+")
	WhitespaceRegex = regexp.MustCompile("\t+|  +")
	PadRegex        = regexp.MustCompile("[{}()\\[\\]<>]")
)

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	sb := strings.Builder{}
	_, err = sb.Write(bytes)
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}

func collapseWhitespace(str string) string {
	// get rid of line endings
	str = LineEndingRegex.ReplaceAllString(str, " ")
	// start collapsing whitespace
	for WhitespaceRegex.MatchString(str) {
		str = WhitespaceRegex.ReplaceAllString(str, " ")
	}

	return str
}
