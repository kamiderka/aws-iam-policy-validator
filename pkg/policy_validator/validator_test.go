package policy_validator

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func Test_IsValidEffect(t *testing.T) {
	tests := []struct {
		have string
		want bool
	}{
		{"Allow", true},
		{"Deny", true},
		{"other", false},
	}

	validate := validator.New()
	validate.RegisterValidation("test-valid-effect", isValidEffect)

	for _, item := range tests {
		err := validate.Var(item.have, "test-valid-effect")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func Test_IsValidVersion(t *testing.T) {
	tests := []struct {
		have string
		want bool
	}{
		{"2012-10-17", true},
		{"2008-10-17", true},
		{"other", false},
	}

	validate := validator.New()
	validate.RegisterValidation("test-valid-version", isValidVersion)

	for _, item := range tests {
		err := validate.Var(item.have, "test-valid-version")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

// func TestIsValidPoliceName(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		valid bool
// 	}{
// 		{"valid-name", true},
// 		{"invalid_name", false},
// 		{"invalid:name", false},
// 	}

// 	for _, tt := range tests {
// 		assert.Equal(t, tt.valid, isValidPoliceName(validator.FieldLevel{Field: validator.Field{Value: tt.name}}))
// 	}
// }

// func TestIsUniqueStatementSids(t *testing.T) {
// 	tests := []struct {
// 		statements []*Statement
// 		valid      bool
// 	}{
// 		{[]*Statement{{Sid: "sid1"}, {Sid: "sid2"}}, true},
// 		{[]*Statement{{Sid: "sid1"}, {Sid: "sid1"}}, false},
// 		{[]*Statement{{Sid: "sid1"}, {Sid: ""}}, true}, // Empty SID is valid
// 	}

// 	for _, tt := range tests {
// 		assert.Equal(t, tt.valid, isUniqueStatementSids(validator.FieldLevel{Field: validator.Field{Value: tt.statements}}))
// 	}
// }

// func TestIsValidAction(t *testing.T) {
// 	tests := []struct {
// 		action string
// 		valid  bool
// 	}{
// 		{"s3:GetObject", true},
// 		{"ec2:RunInstances", true},
// 		{"invalid", false},
// 	}

// 	for _, tt := range tests {
// 		assert.Equal(t, tt.valid, isValidAction(validator.FieldLevel{Field: validator.Field{Value: tt.action}}))
// 	}
// }

// func TestIsValidSid(t *testing.T) {
// 	tests := []struct {
// 		sid   string
// 		valid bool
// 	}{
// 		{"sid1", true},
// 		{"sid-1", true},
// 		{"invalid:sid", false},
// 	}

// 	for _, tt := range tests {
// 		assert.Equal(t, tt.valid, isValidSid(validator.FieldLevel{Field: validator.Field{Value: tt.sid}}))
// 	}
// }

// func TestNewValidator(t *testing.T) {
// 	v := NewValidator()
// 	require.NotNil(t, v)

// 	// Test whether the custom validations are registered
// 	_, ok := v.TagMap()["valid-effect"]
// 	assert.True(t, ok)
// 	_, ok = v.TagMap()["valid-version"]
// 	assert.True(t, ok)
// 	_, ok = v.TagMap()["valid-police-name"]
// 	assert.True(t, ok)
// 	_, ok = v.TagMap()["unique-sids"]
// 	assert.True(t, ok)
// 	_, ok = v.TagMap()["valid-action"]
// 	assert.True(t, ok)
// 	_, ok = v.TagMap()["valid-sid"]
// 	assert.True(t, ok)
// }
