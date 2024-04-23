package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filePath := flag.String("file", "", "Ścieżka do pliku JSON do walidacji")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Nie podano ścieżki do pliku JSON")
		os.Exit(1)
	}

	jsonFile, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println("Błąd podczas odczytu pliku:", err)
		os.Exit(1)
	}

	var policy AwsIamRolePolicy
	if err := json.Unmarshal(jsonFile, &policy); err != nil {
		fmt.Println("Błąd podczas deserializacji danych JSON:", err)
		return
	}

	validate = NewValidator()
	if err = validate.Struct(policy); err != nil {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}