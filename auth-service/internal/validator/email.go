package validator

func (v *Validator) IsEmailValid(email string) bool {
	return v.emailValidateRegexp.MatchString(email)
}
