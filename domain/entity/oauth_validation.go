package entity

import validation "github.com/go-ozzo/ozzo-validation/v4"

func (e *RequestToken) Validate() error {
	var rules []*validation.FieldRules
	rules = append(rules,
		validation.Field(&e.GrantType, validation.Required),
		validation.Field(&e.ClientID, validation.Required),
		validation.Field(&e.ClientSecret, validation.Required),
	)

	if e.GrantType == RefreshToken {
		rules = append(rules, validation.Field(&e.RefreshToken, validation.Required))
	}

	return validation.ValidateStruct(e, rules...)
}
