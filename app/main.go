package main

import (
	"fmt"
	"log-pipeline/generator"
	"os"
)

func main() {
	output, err := makeOutput("output.log")
	if err != nil {
		panic(err)
	}

	logger := generator.NewLogger(generator.ApacheCommonLine, output, 1)
	fmt.Println("Logging Start")
	err = logger.GenerateLogs()
	if err != nil {
		panic(err)
	}
}

func makeOutput(path string) (*os.File, error) {
	if path == "" {
		return os.Stdout, nil
	}
	return os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
}
