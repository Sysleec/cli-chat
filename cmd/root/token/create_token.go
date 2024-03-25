package token

import (
	"fmt"
	"os"
)

func CreateRefresh(refreshToken string) error {
	file, err := os.Create("bin/refreshtoken.txt")
	if err != nil {
		return fmt.Errorf("failed to create or open file: %v", err)
	}
	defer file.Close()
	_, err = file.Write([]byte(refreshToken))
	if err != nil {
		return fmt.Errorf("failed to write in file: %v", err)
	}

	return nil
}

func CreateAccess(accessToken string) error {
	file, err := os.Create("bin/accesstoken.txt")
	if err != nil {
		return fmt.Errorf("failed to create or open file: %v", err)
	}
	defer file.Close()
	_, err = file.Write([]byte(accessToken))
	if err != nil {
		return fmt.Errorf("failed to write in file: %v", err)
	}

	return nil
}
