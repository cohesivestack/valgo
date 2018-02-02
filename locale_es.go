package valgo

func init() {
	setDefaultSpanishMessages()
}

func setDefaultSpanishMessages() {
	getLocales()["es"] = locale{
		Messages: map[string]string{
			"blank":            "\"{{Title}}\" debe estar en blanco",
			"not_blank":        "\"{{Title}}\" no puede estar en blanco",
			"empty":            "\"{{Title}}\" debe estar vacío",
			"not_empty":        "\"{{Title}}\" no puede estar vacío",
			"equal_to":         "\"{{Title}}\" debe ser igual a \"{{Value}}\"",
			"not_equal_to":     "\"{{Title}}\" no puede ser igual a \"{{Value}}\"",
			"identical_to":     "\"{{Title}}\" debe ser igual a \"{{Value}}\"",
			"not_identical_to": "\"{{Title}}\" no puede ser igual a \"{{Value}}\"",
			"a_number":         "\"{{Title}}\" debe ser un número",
			"an_integer":       "\"{{Title}}\" debe ser un número entero",
		},
	}
}
