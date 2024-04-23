package policy_validator

type AwsIamRolePolicy struct {
	PolicyName     string         `json:"PolicyName" validate:"required,max=128,valid-police-name"`
	PolicyDocument *PolicyDocument `json:"PolicyDocument" validate:"required"`
}

type PolicyDocument struct {
	Version   string      `json:"Version" validate:"required,valid-version" `
	Statement []*Statement `json:"Statement" validate:"unique-sids,dive"`
}
type Statement struct {
	Sid      	string   	`json:"Sid,omitempty" validate:"valid-sid"`
	Effect   	string   	`json:"Effect" validate:"required,valid-effect"`
	Action   	[]string 	`json:"Action" validate:"required,dive,valid-action"`
	Resource 	string   	`json:"Resource" validate:"required,ne=*"`
}