package valgo

func setDefaultEnglishMessages() {
	getLocales()["en"] = &locale{
		Messages: map[string]string{
			"valid":                   "\"{{value}}\" is not a valid value for {{title}}",
			"not_valid":               "\"{{value}}\" is not a valid value for {{title}}",
			"blank":                   "{{title}} must be blank",
			"not_blank":               "{{title}} can't be blank",
			"empty":                   "{{title}} must be empty",
			"not_empty":               "{{title}} can't be empty",
			"equivalent_to":           "{{title}} must be equal to \"{{value}}\"",
			"not_equivalent_to":       "{{title}} can't be equal to \"{{value}}\"",
			"equal_to":                "{{title}} must be equal to \"{{value}}\"",
			"not_equal_to":            "{{title}} can't be equal to \"{{value}}\"",
			"greater_than":            "{{title}} must be greater than \"{{value}}\"",
			"not_greater_than":        "{{title}} can't be greater than \"{{value}}\"",
			"greater_or_equal_to":     "{{title}} must be greater or equal to \"{{value}}\"",
			"not_greater_or_equal_to": "{{title}} can't be greater or equal to \"{{value}}\"",
			"less_than":               "{{title}} must be less than \"{{value}}\"",
			"not_less_than":           "{{title}} can't be less than \"{{value}}\"",
			"less_or_equal_to":        "{{title}} must be less or equal to \"{{value}}\"",
			"not_less_or_equal_to":    "{{title}} can't be less or equal to \"{{value}}\"",
			"identical_to":            "{{title}} must be equal to \"{{value}}\"",
			"not_identical_to":        "{{title}} can't be equal to \"{{value}}\"",
			"matching_to":             "{{title}} must match to \"{{value}}\"",
			"not_matching_to":         "{{title}} can't match to \"{{value}}\"",
			"a_number":                "{{title}} must be a number",
			"a_number_type":           "{{title}} must be a number type",
			"an_integer":              "{{title}} must be an integer number",
			"an_integer_type":         "{{title}} must be an integer type",
			"a_string":                "{{title}} must be a text",
			"an_email":                "{{title}} is not an email address",
			"not_an_email":            "{{title}} must not be an email address",
			"in_slice":                "\"{{value}}\" is not a valid value for {{title}}",
			"not_in_slice":            "\"{{value}}\" is not a valid value for {{title}}",
		},
	}
}
