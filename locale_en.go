package valgo

func setDefaultEnglishMessages() {
	getLocales()["en"] = &locale{
		Messages: map[string]string{
			"blank":                   "{{title}} must be blank",
			"not_blank":               "{{title}} can't be blank",
			"empty":                   "{{title}} must be empty",
			"not_empty":               "{{title}} can't be empty",
			"equivalent_to":           "{{title}} must be equal to \"{{Value}}\"",
			"not_equivalent_to":       "{{title}} can't be equal to \"{{Value}}\"",
			"equal_to":                "{{title}} must be equal to \"{{Value}}\"",
			"not_equal_to":            "{{title}} can't be equal to \"{{Value}}\"",
			"greater_than":            "{{title}} must be greater than \"{{Value}}\"",
			"not_greater_than":        "{{title}} can't be greater than \"{{Value}}\"",
			"greater_or_equal_to":     "{{title}} must be greater or equal to \"{{Value}}\"",
			"not_greater_or_equal_to": "{{title}} can't be greater or equal to \"{{Value}}\"",
			"less_than":               "{{title}} must be less than \"{{Value}}\"",
			"not_less_than":           "{{title}} can't be less than \"{{Value}}\"",
			"less_or_equal_to":        "{{title}} must be less or equal to \"{{Value}}\"",
			"not_less_or_equal_to":    "{{title}} can't be less or equal to \"{{Value}}\"",
			"identical_to":            "{{title}} must be equal to \"{{Value}}\"",
			"not_identical_to":        "{{title}} can't be equal to \"{{Value}}\"",
			"matching_to":             "{{title}} must match to \"{{Value}}\"",
			"not_matching_to":         "{{title}} can't match to \"{{Value}}\"",
			"a_number":                "{{title}} must be a number",
			"a_number_type":           "{{title}} must be a number type",
			"an_integer":              "{{title}} must be an integer number",
			"an_integer_type":         "{{title}} must be an integer type",
			"a_string":                "{{title}} must be a text",
			"an_email":                "{{title}} is not an email address",
			"not_an_email":            "{{title}} must not be an email address",
		},
	}
}
