package main

import (
	// "github.com/deckarep/golang-set"

	"github.com/go-playground/validator/v10"
	"github.com/golang-collections/collections/set"
)

func NewValidator() *validator.Validate {
	validator := validator.New()

	validator.RegisterValidation("valid-effect", isValidEffect)
	validator.RegisterValidation("valid-version", isValidVersion)
	validator.RegisterValidation("valid-police-name", isValidPoliceName)
	validator.RegisterValidation("unique-sids", isUniqueStatementSids)

	return validator
}

func isValidEffect(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case
		StatementEffectDeny,
		StatementEffectAllowed:
		return true
	default:
		return false
	}
}

func isValidVersion(fl validator.FieldLevel) bool {
	switch fl.Field().String() {
	case
		PolicyDocumentVersion2012,
		PolicyDocumentVersion2008:
		return true
	default:
		return false
	}
}

func isValidPoliceName(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	return policyNameRegex.MatchString(field)
}

func isUniqueStatementSids(fl validator.FieldLevel) bool {
	field := fl.Field()
	slice, ok := field.Interface().([]*Statement)
	if !ok {
		return false
	}

	prev := set.New()

	for _, element := range slice {

		if prev.Has(element.Sid) {
			return false
		}

		if element.Sid != "" {
			prev.Insert(element.Sid)

		}
	}
	return true
}

