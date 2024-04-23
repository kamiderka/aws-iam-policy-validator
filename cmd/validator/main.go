package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kamiderka/aws-iam-policy-validator/pkg/policy_validator"
)

func main() {
	filePath := flag.String("file", "", "Path to the JSON file to validate")
	dirPath := flag.String("dir", "", "Path to the directory containing JSON files to validate")
	flag.Parse()

	if *filePath != "" && *dirPath == "" {
		if result, err := validateSingleFile(*filePath); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		} else {
			fmt.Printf("%s | %t\n", filepath.Base(*filePath), result)
		}
		return
	}

	if *dirPath != "" && *filePath == "" {
		if err := validateDirectory(*dirPath); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		return
	}

	fmt.Println("Usage: ")
	flag.PrintDefaults()
}

func validateSingleFile(filePath string) (bool, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("error opening file: %v", err)
	}
	defer jsonFile.Close()

	var policy policy_validator.AwsIamRolePolicy
	if err := json.NewDecoder(jsonFile).Decode(&policy); err != nil {
		return false, fmt.Errorf("error deserializing JSON data: %v", err)
	}

	validate := policy_validator.NewValidator()
	if err := validate.Struct(policy); err != nil {
		return false, nil
	}

	return true, nil
}

func validateDirectory(dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return fmt.Errorf("error opening directory: %v", err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(dirPath, file.Name())
			if result, err := validateSingleFile(filePath); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("%s | %t\n", file.Name(), result)
			}
		}
	}

	return nil
}
