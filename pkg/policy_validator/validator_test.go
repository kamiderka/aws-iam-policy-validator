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

func Test_IsValidPoliceName(t *testing.T) {
	tests := []struct {
		have  string
		want bool
	}{
		{"validName", true},            
		{"Valid_Name", true},          
		{"ValidName123", true},        
		{"123ValidName", true},        
		{"Valid-Name", true},          
		{"Valid.Name", true},           
		{"Valid@Name", true},           
		{"Valid=Name", true},           
		{"Valid,Name", true},           
		{"Valid=Name,123", true},       
		{"Valid_Name@123", true},       
		{"Valid_Name@", true},       
		{"valid-Name,", true},
		{"Invalid-Name!", false},  
		{"", false},                    
		{" ", false},                   
		{"Invalid Name", false},        
		{"Invalid\nName", false},
		{"Invalid:Name", false},        
	}  

		validate := validator.New()
		validate.RegisterValidation("test-valid-police-name", isValidPoliceName)
	
		for _, item := range tests {
			err := validate.Var(item.have, "test-valid-police-name")
			if item.want {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		}
	}

func TestIsUniqueStatementSids(t *testing.T) {
	tests := []struct {
		have 	[]*Statement
		want      bool
	}{
		{[]*Statement{{Sid: "sid1"}, {Sid: "sid2"}}, true},

		{[]*Statement{{Sid: "sid1"}, {Sid: "sid1"}}, false},
	
		{[]*Statement{{Sid: "sid1"}, {Sid: ""}}, true}, 
	
		{[]*Statement{{Sid: ""}, {Sid: ""}}, true},
		 
		{[]*Statement{{Sid: "sid1"}, {Sid: "sid2"},
					{Sid: "sid3"}, {Sid: "sid4"}}, true},

		{[]*Statement{{Sid: "sid1"}, {Sid: "sid2"},
					{Sid: "sid3"}, {Sid: "sid1"}}, false},

		{[]*Statement{{}}, true},

		}

		validate := validator.New()
		validate.RegisterValidation("test-unique-statement-sids", isUniqueStatementSids)
	
		for _, item := range tests {
			err := validate.Var(item.have, "test-unique-statement-sids")
			if item.want {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		}
}

func TestIsValidAction(t *testing.T) {
	tests := []struct {
		have  string
		want  bool
	}{
		{"s3:GetObject", true},
		{"s3:PutObject", true},
		{"ec2:DescribeInstances", true},
		{"dynamodb:PutItem", true},
		{"InvalidAction", false},            
		{"s3:GetObject*", false},
		{"s3:get_object*", false},            
		{"ec2:*", true},                    
		{"dynamodb:PutItem?", false},  
		{"s3:GetObjA2:s", false}, 
		{"s3::GetObjA2", false},      
		{"s3:GetObject, s3:PutObject", false}, 
		{"s3:GetObject,InvalidAction", false}, 
		{"s3:GetObject InvalidAction", false}, 
		{"", false},                         
	}
	
	validate := validator.New()
	validate.RegisterValidation("test-valid-action", isValidAction)

	for _, item := range tests {
		err := validate.Var(item.have, "test-valid-action")
		if item.want {
			assert.Nil(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestIsValidSid(t *testing.T) {
	tests := []struct {
		have   	string
		want 	bool
	}{
		{"sid1", true},           
		{"sid_2", false},          
		{"s3", true},             
		{"invalid!sid", false},   
		{"s-id,sid", false},      
		{"", true},
		{" ", false},   
		{"s-id sid", false},      
		{"s-id'sid", false},     
		{"s-id&sid", false},      
		{"s-id@", false},        
		{"sid123456789012345678901234567890123456789012345678901234567890", true}, 
		}

		validate := validator.New()
		validate.RegisterValidation("test-valid-sid", isValidSid)
	
		for _, item := range tests {
			err := validate.Var(item.have, "test-valid-sid")
			if item.want {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err)
			}
		}
		}
