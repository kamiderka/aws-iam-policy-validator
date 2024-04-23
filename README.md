<h1 align="center">
AWS IAM Policy Validator
</h1>
A simple command-line tool written in Go that allows you to validate AWS Identity and Access Management (IAM) policies in JSON format. It helps you ensure that your policies adhere to the required structure and syntax.


## Getting started 
```go
go install -v github.com/kamiderka/aws-iam-policy-validator/cmd/validator@latest`
```
# Usage
```sh
validator
```
This will display help for the tool. Here are all the switches it supports.

```console
Usage:
  validator [flags]

Flags:
INPUT:
  -dir string
        Path to the directory containing JSON files to validate
  -file string
        Path to the JSON file to validate
```

## Output
The tool will output the result of the validation for each file in the format:

```console
filename.json | true/false
```
>Note: The decision to consider a file invalid if the Resource field in the input JSON contains a single asterisk (*) is solely based on the client's request.

