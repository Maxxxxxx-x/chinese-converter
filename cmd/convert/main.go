package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/longbridgeapp/opencc"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: ./convert [OPTIONS]\n")
	flag.PrintDefaults()
}

func readFile(filePath string) string {
	body, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}

	return string(body)
}

func getFileNameAndExtension(pathName string) (string, string) {
    fileExt := filepath.Ext(pathName)
    fileName := strings.Split(pathName, fileExt)[0]
    return fileName, fileExt
}

func convert(content string) string {
	s2hk, err := opencc.New("s2hk")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start opencc.\nErr: %v", err)
		os.Exit(1)
	}

	output, err := s2hk.Convert(content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to convert text.\nErr: %v\n", err)
		os.Exit(1)
	}

	return output
}

func writeToFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Failed to create file %s. Err: %v\n", fileName, err)
		os.Exit(1)
	}

	defer file.Close()

	file.WriteString(content)
}

func main() {
	var (
		inputFilePath  string
		outputFileName string
	)

	flag.StringVar(&inputFilePath, "i", "", "[REQUIRED] path to input file")
	flag.StringVar(&outputFileName, "o", "", "[OPTIONAL] name of the output file")
	flag.Parse()

	if inputFilePath == "" {
		fmt.Fprintf(os.Stderr, "Missing input file path\n")
		usage()
		os.Exit(1)
	}
	if _, err := os.Stat(inputFilePath); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "File %s does not exist!\n", inputFilePath)
		os.Exit(1)
	}

	fmt.Println("Starting conversion...")
	fileName, fileExt := getFileNameAndExtension(inputFilePath)
	if outputFileName == "" {
		outputFileName = fmt.Sprintf("%s.zh_HK.%s", fileName, fileExt)
	}
	fileContent := readFile(inputFilePath)

	output := convert(fileContent)
	fmt.Printf("Converted!\nCreating file %s...\n", outputFileName)
	writeToFile(outputFileName, output)
	fmt.Println("Done!")
}
