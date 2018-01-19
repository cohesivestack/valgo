package valgo

func init() {
	setDefaultEnglishMessages()
}

func setDefaultEnglishMessages() {
	getLocales()["en"] = locale{
		Messages: map[string]string{
			"blank":     "\"{{Title}}\" must be blank",
			"not_blank": "\"{{Title}}\" can't be blank",
			"empty":     "\"{{Title}}\" must be empty",
			"not_empty": "\"{{Title}}\" can't be empty",
		},
	}
}
