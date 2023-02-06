package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

//creating the functions to check the validity of the form inputs

// return valid as true if the form has no errors
func (f *Form) Valid() bool { return len(f.Errors) == 0 }

func CreateForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// takes multiple arguments
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		//this checks if the input of the form field is empty (trim space removes unnecessary whitespace)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "this field cannot be empty")
		}
	}
}

// if the input is less than the required return error
func (f *Form) MinLength(field string, minDesiredLength int, Req *http.Request) bool {
	x := Req.Form.Get(field)

	if len(x) < minDesiredLength {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters", minDesiredLength))
		return false
	}
	return true
}

func (f *Form) IsValidEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid Email Adress")
	}
}

// func (f *Form) IsValidPhoneNumber(Numberfield *phonenumbers.PhoneNumber, Regionfield string) bool {
// 	return phonenumbers.IsValidNumberForRegion(Numberfield, Regionfield)
// }

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	} else {
		return true
	}
}
