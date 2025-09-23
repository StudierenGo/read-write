package files

import (
	"fmt"
	"os"
)

func ReadFile(fileName string) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:", string(data))
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
