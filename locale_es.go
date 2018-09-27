package valgo

func setDefaultSpanishMessages() {
	getLocales()["es"] = &locale{
		Messages: map[string]string{
			"blank":             "{{title}} debe estar en blanco",
			"not_blank":         "{{title}} no puede estar en blanco",
			"empty":             "{{title}} debe estar vacío",
			"not_empty":         "{{title}} no puede estar vacío",
			"equivalent_to":     "{{title}} debe ser igual a \"{{value}}\"",
			"not_equivalent_to": "{{title}} no puede ser igual a \"{{value}}\"",
			"equal_to":          "{{title}} debe ser igual a \"{{value}}\"",
			"not_equal_to":      "{{title}} no puede ser igual a \"{{value}}\"",
			"matching_to":       "{{title}} debe corresponder a \"{{value}}\"",
			"not_matching_to":   "{{title}} no puede corresponder a \"{{value}}\"",
			"identical_to":      "{{title}} debe ser igual a \"{{value}}\"",
			"not_identical_to":  "{{title}} no puede ser igual a \"{{value}}\"",
			"a_number":          "{{title}} debe ser un número",
			"a_number_type":     "{{title}} debe ser un tipo número",
			"an_integer":        "{{title}} debe ser un número entero",
			"an_integer_type":   "{{title}} debe ser un número tipo entero",
			"a_string":          "{{title}} debe ser un texto",
			"an_email":          "{{title}} no es una dirección de correo valida",
			"not_an_email":      "{{title}} debe ser una dirección de correo",
		},
	}
}
