package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		// fmt.Println("Error reading file:", err)
		return nil, err
	}

	return data, nil
}

func WriteFile(name string, content []byte) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println("Error creating file:", err)
	}

	_, err = file.Write(content)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File written successfully!")
	defer file.Close()
}
