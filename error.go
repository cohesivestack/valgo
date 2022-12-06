package valgo

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasttemplate"
)

// Error Implementation of the Go error interface in Valgo. The [Validation.Error()]
// method returns a value of this type.
//
// There is a function in this type, [Errors()], that returns a list of errors
// in a [Validation] session.
type Error struct {
	errors map[string]*valueError
}

type errorTemplate struct {
	key      string
	template *string
	params   map[string]interface{}
}

// valueError Contains information about each invalid field value returned by the
// [Validation] session.
type valueError struct {
	name           *string
	title          *string
	errorTemplates map[string]*errorTemplate
	errorMessages  []string
	messages       []string
	dirty          bool
	validator      *Validation
}

// Title The title of the invalid field value.
func (ve *valueError) Title() string {
	return *ve.title
}

// Name The name of the invalid field value.
func (ve *valueError) Name() string {
	return *ve.name
}

// Messages Error messages related to an invalid field value.
func (ve *valueError) Messages() []string {
	if ve.dirty {
		ve.messages = []string{}
		for _, et := range ve.errorTemplates {
			ve.messages = append(ve.messages, ve.buildMessageFromTemplate(et))
		}

		ve.messages = append(ve.messages, ve.errorMessages...)

		ve.dirty = false
	}

	return ve.messages
}

func (ve *valueError) buildMessageFromTemplate(et *errorTemplate) string {
	var ts string
	if et.template != nil {
		ts = *et.template
	} else if _ts, ok := ve.validator._locale.Messages[et.key]; ok {
		ts = _ts
	} else {
		ts = concatString("ERROR: THERE IS NOT A MESSAGE WITH THE KEY: ", et.key)
	}

	var title string
	if ve.title == nil {
		title = humanizeName(*ve.name)
	} else {
		title = *ve.title
	}

	et.params["name"] = *ve.name
	et.params["title"] = title

	t := fasttemplate.New(ts, "{{", "}}")

	// Ensure interface{} values are string in order to be handle by fasttemplate
	for k, v := range et.params {
		if k != "name" && k != "title" {
			et.params[k] = fmt.Sprintf("%v", v)
		}
	}

	return t.ExecuteString(et.params)
}

// Error Return the error message associated with a Valgo error.
func (e *Error) Error() string {
	count := len(e.errors)
	if count == 1 {
		return "There is 1 error"
	}

	return fmt.Sprintf("There are %v errors", count)
}

// Errors Return a map with each Invalid value error.
//
//nolint:revive // by design. should be exported as can be annoying to use
func (e *Error) Errors() map[string]*valueError {
	return e.errors
}

// MarshalJSON Return the JSON encoding of the validation error messages.
//
// A custom function can be set with [SetMarshalJson()].
func (e *Error) MarshalJSON() ([]byte, error) {
	if customMarshalJSON != nil {
		return customMarshalJSON(e)
	}

	errors := map[string]interface{}{}

	for k, v := range e.errors {
		errors[k] = v.Messages()
	}

	return json.Marshal(errors)
}
