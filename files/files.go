package files

import (
	"os"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFile(name string, content []byte) {
	file, err := os.Create(name)

	if err != nil {
		color.Red("Error creating file:", err)
	}

	_, err = file.Write(content)

	if err != nil {
		color.Red("Error writing to file:", err)
		return
	}

	defer file.Close()
}
