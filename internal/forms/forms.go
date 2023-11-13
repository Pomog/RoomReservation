package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	Errors errors
	url.Values
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
	return x != ""
}
