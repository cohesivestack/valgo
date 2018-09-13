package valgo

func init() {
	setDefaultSpanishMessages()
}

func setDefaultSpanishMessages() {
	getLocales()["es"] = locale{
		Messages: map[string]string{
			"blank":             "\"{{Title}}\" debe estar en blanco",
			"not_blank":         "\"{{Title}}\" no puede estar en blanco",
			"empty":             "\"{{Title}}\" debe estar vacío",
			"not_empty":         "\"{{Title}}\" no puede estar vacío",
			"equivalent_to":     "\"{{Title}}\" debe ser igual a \"{{Value}}\"",
			"not_equivalent_to": "\"{{Title}}\" no puede ser igual a \"{{Value}}\"",
			"equal_to":          "\"{{Title}}\" debe ser igual a \"{{Value}}\"",
			"not_equal_to":      "\"{{Title}}\" no puede ser igual a \"{{Value}}\"",
			"matching_to":       "\"{{Title}}\" debe corresponder a \"{{Value}}\"",
			"not_matching_to":   "\"{{Title}}\" no puede corresponder a \"{{Value}}\"",
			"identical_to":      "\"{{Title}}\" debe ser igual a \"{{Value}}\"",
			"not_identical_to":  "\"{{Title}}\" no puede ser igual a \"{{Value}}\"",
			"a_number":          "\"{{Title}}\" debe ser un número",
			"a_number_type":     "\"{{Title}}\" debe ser un tipo número",
			"an_integer":        "\"{{Title}}\" debe ser un número entero",
			"an_integer_type":   "\"{{Title}}\" debe ser un número tipo entero",
			"a_string":          "\"{{Title}}\" debe ser un texto",
			"an_email":          "\"{{Title}}\" no es una dirección de correo valida",
			"not_an_email":      "\"{{Title}}\" debe ser una dirección de correo",
		},
	}
}
