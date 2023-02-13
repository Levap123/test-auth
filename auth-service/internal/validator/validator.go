package validator

import "regexp"

type Validator struct {
	emailValidateRegexp *regexp.Regexp
}

func NewValidator() *Validator {
	return &Validator{
		emailValidateRegexp: regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"),
	}
}
