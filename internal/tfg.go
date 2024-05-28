package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func tfgExists(projectDir string) (bool, error) {
	_, err := os.Stat(projectDir)

	if os.IsNotExist(err) {
		return false, nil
	}

	_, err = os.Stat(filepath.Join(projectDir, TFG_FILENAME))
	if os.IsNotExist(err) {
		return false, fmt.Errorf("a folder with the same name already exists")
	}

	return true, nil
}
