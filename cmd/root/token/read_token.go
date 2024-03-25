package token

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadRefresh() (string, error) {
	return readFile("bin/refreshtoken.txt")
}

func ReadAccess() (string, error) {
	return readFile("bin/accesstoken.txt")
}

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read from file: %v", err)
	}

	// line, err = strconv.Unquote(line)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to unquote file string: %v", err)
	// }
	return line, nil
}
