# AWS IAM Policy Validator 
(Rudimentary README, it will be better in few hours)
## Description
The AWS IAM Policy Validator is a simple command-line tool written in Go that allows you to validate AWS Identity and Access Management (IAM) policies in JSON format. It helps you ensure that your policies adhere to the required structure and syntax.

## Getting Started

```
git clone https://github.com/kamiderka/aws-iam-policy-validator.git
cd ./aws-iam-policy-validator
```

From this place you can: 

- use binary: `./bin/validator.exe`
- build it by yourself:  `go build cmd/validator .`
- or run it by yourself `go run cmd/validator/main.go`

## Usage:
````
  -dir string
        Path to the directory containing JSON files to validate
  -file string
        Path to the JSON file to validate
````
Makefile coming soon! 
