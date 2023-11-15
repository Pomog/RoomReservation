package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	Errors errors
	url.Values
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form structure
func New(data url.Values) *Form {
	return &Form{
		Errors: make(errors),
		Values: data,
	}
}

// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// Required checks if specific fields in form data are present and not blank
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks for string minimum length
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field is too short must be at least %d characters", length))
		return false
	}
	return true
}
