package main

import "regexp"

const (
	// policyNameRegexString = "[\\w+=,.@-]+"
	policyNameRegexString = "^[a-zA-Z0-9+=,.@_-]+$"
)

var (
	policyNameRegex = regexp.MustCompile(policyNameRegexString)
)