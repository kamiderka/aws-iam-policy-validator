<h1 align="center">
AWS IAM Policy Validator
</h1>

<p align="center">
  <a href="#description">Description</a> •
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#input">Input</a> •
  <a href="#output">Output</a> 
</p>

# Description 
A simple command-line tool written in Go that allows you to validate AWS Identity and Access Management (IAM) policies in JSON format. It helps you ensure that your policies adhere to the required structure and syntax.

# Installation 
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

# Input
An example of an input JSON object representing an IAM policy:
```json
{
    "PolicyName": "root",
    "PolicyDocument": {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "IamListAccess",
                "Effect": "Allow",
                "Action": [
                    "iam:ListRoles",
                    "iam:ListUsers"
                ],
                "Resource": "*"
            }
        ]
    }
}
```
>Note: Inside the `test` directory, you'll find examples of both valid and invalid inputs that you can use for validation.

# Output
The tool will output the result of the validation for each file in the format:

```console
filename.json | true/false
```
>Note: The decision to consider a file invalid if the Resource field in the input JSON contains a single asterisk (*) is solely based on the client's request.

