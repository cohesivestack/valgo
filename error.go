package valgo

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasttemplate"
)

type Error struct {
	errors map[string]*valueError
}

type errorTemplate struct {
	key      string
	template *string
	values   map[string]interface{}
}

type valueError struct {
	name           *string
	title          *string
	errorTemplates map[string]*errorTemplate
	errorMessages  []string
	messages       []string
	dirty          bool
	validator      *Validator
}

func (ve *valueError) Title() string {
	return *ve.title
}

func (ve *valueError) Name() string {
	return *ve.name
}

func (ve *valueError) Messages() []string {
	if ve.dirty {
		ve.messages = []string{}
		for _, et := range ve.errorTemplates {
			ve.messages = append(ve.messages, ve.buildMessageFromTemplate(et))
		}
		for _, em := range ve.errorMessages {
			ve.messages = append(ve.messages, em)
		}
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

	et.values["name"] = *ve.name
	et.values["title"] = title

	t := fasttemplate.New(ts, "{{", "}}")

	// Ensure interface{} values are string in order to be handle by fasttemplate
	for k, v := range et.values {
		if k != "name" && k != "title" {
			et.values[k] = fmt.Sprintf("%v", v)
		}
	}

	return t.ExecuteString(et.values)
}

func (e *Error) Error() string {
	count := len(e.errors)
	if count == 1 {
		return fmt.Sprintf("There is 1 error")
	} else {
		return fmt.Sprintf("There are %v errors", count)
	}
}

func (e *Error) Errors() map[string]*valueError {
	return e.errors
}

func (e *Error) MarshalJSON() ([]byte, error) {
	errors := map[string]interface{}{}

	for k, v := range e.errors {
		if len(v.Messages()) == 1 {
			errors[k] = v.Messages()[0]
		} else {
			errors[k] = v.Messages()
		}
	}

	return json.Marshal(errors)
}
