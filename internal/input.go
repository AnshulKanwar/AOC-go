package internal

import (
	"log"
	"os"
)

func ReadInputFile(filepath string) string {
	input, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}
