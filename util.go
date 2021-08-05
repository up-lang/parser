package parser

import (
	"io/ioutil"
	"os"
	"strings"
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
