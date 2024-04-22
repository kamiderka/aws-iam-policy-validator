package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	// Odczytanie danych JSON z pliku
	jsonBytes, err := ioutil.ReadFile("wrong_policy.json")
	if err != nil {
		fmt.Println("Błąd podczas odczytu pliku:", err)
		return
	}

	// Zmapowanie danych JSON na strukturę AwsIamRolePolicy
	var policy AwsIamRolePolicy
	if err := json.Unmarshal(jsonBytes, &policy); err != nil {
		fmt.Println("Błąd podczas deserializacji danych JSON:", err)
		return
	}

	// Wyświetlenie odczytanej polityki
	fmt.Println("PolicyName:", policy.PolicyName)
	fmt.Println("Version:", policy.PolicyDocument.Version)
	fmt.Println("Statements:")
	for _, stmt := range policy.PolicyDocument.Statement {
		fmt.Println("  Sid:", stmt.Sid)
		fmt.Println("  Effect:", stmt.Effect)
		fmt.Println("  Action:", stmt.Action)
		fmt.Println("  Resource:", stmt.Resource)
	}
	validate = NewValidator()
	validate.RegisterValidation("valid-effect", isValidEffect)
	if err = validate.Struct(policy); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("end ")
}
