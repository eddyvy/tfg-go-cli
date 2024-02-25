package helloworld

import (
	"fmt"
	"os"
)

const DIRECTORY_PERMISSION = 0755
const HELLO_WORLD_GO_TEXT = "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Hello, World!\")\n}"

func CreateHelloWorld(path string) error {
	// Create a new directory if it doesn't exist
	err := os.Mkdir(path, DIRECTORY_PERMISSION)
	if err != nil {
		fmt.Printf("Error creating directory, check that \"%s\" directory doesn't exist\n", path)
		return err
	}

	// Create a new file
	file, err := os.Create(path + "/main.go")
	if err != nil {
		return err
	}
	defer file.Close()

	// Write to the file
	_, err = file.WriteString(HELLO_WORLD_GO_TEXT)
	if err != nil {
		return err
	}

	return nil
}
