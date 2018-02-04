package valgo

func init() {
	setDefaultEnglishMessages()
}

func setDefaultEnglishMessages() {
	getLocales()["en"] = locale{
		Messages: map[string]string{
			"blank":             "\"{{Title}}\" must be blank",
			"not_blank":         "\"{{Title}}\" can't be blank",
			"empty":             "\"{{Title}}\" must be empty",
			"not_empty":         "\"{{Title}}\" can't be empty",
			"equivalent_to":     "\"{{Title}}\" must be equal to \"{{Value}}\"",
			"not_equivalent_to": "\"{{Title}}\" can't be equal to \"{{Value}}\"",
			"equal_to":          "\"{{Title}}\" must be equal to \"{{Value}}\"",
			"not_equal_to":      "\"{{Title}}\" can't be equal to \"{{Value}}\"",
			"identical_to":      "\"{{Title}}\" must be equal to \"{{Value}}\"",
			"not_identical_to":  "\"{{Title}}\" can't be equal to \"{{Value}}\"",
			"a_number":          "\"{{Title}}\" must be a number",
			"a_number_type":     "\"{{Title}}\" must be a number type",
			"an_integer":        "\"{{Title}}\" must be an integer number",
			"an_integer_type":   "\"{{Title}}\" must be an integer type",
			"a_string":          "\"{{Title}}\" must be a text",
		},
	}
}
