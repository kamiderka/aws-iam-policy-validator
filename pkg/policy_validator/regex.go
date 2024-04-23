package policy_validator

import "regexp"

const (
	policyNameRegexString 		= "^[\\w+=,.@-]+$"
	statementActionRegexString	= "^([a-zA-Z0-9]+):{1}([a-zA-Z]+)$"
	statementSidRegexString		= "^[a-zA-Z0-9]*$"
)

var (
	policyNameRegex 		= regexp.MustCompile(policyNameRegexString)
	statementActionRegex 	= regexp.MustCompile(statementActionRegexString)
	statementSidRegex 		= regexp.MustCompile(statementSidRegexString) 
)