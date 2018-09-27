package valgo

import "github.com/valyala/fasttemplate"

type errorMessage struct {
	key      string
	template *string
	message  string
	values   map[string]interface{}
}

type valueError struct {
	name          *string
	title         *string
	errorMessages map[string]*errorMessage
	messages      []string
	dirty         bool
	validator     *Validator
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
		for _, em := range ve.errorMessages {
			ve.messages = append(ve.messages, ve.buildMessage(em))
		}
		ve.dirty = false
	}

	return ve.messages
}

func (ve *valueError) buildMessage(em *errorMessage) string {

	var ts string
	if em.template != nil {
		ts = *em.template
	} else if _ts, ok := ve.validator._locale.Messages[em.key]; ok {
		ts = _ts
	} else {
		ts = concatString("ERROR: THERE IS NOT A MESSAGE WITH THE KEY: ", em.key)
	}

	var title string
	if ve.title == nil {
		title = humanizeName(*ve.name)
	} else {
		title = *ve.title
	}

	em.values["name"] = *ve.name
	em.values["title"] = title

	t := fasttemplate.New(ts, "{{", "}}")

	return t.ExecuteString(em.values)
}
