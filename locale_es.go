package valgo

func setDefaultSpanishMessages() {
	getLocales()["es"] = &locale{
		Messages: map[string]string{
			"blank":             "{{title}} debe estar en blanco",
			"not_blank":         "{{title}} no puede estar en blanco",
			"empty":             "{{title}} debe estar vacío",
			"not_empty":         "{{title}} no puede estar vacío",
			"equivalent_to":     "{{title}} debe ser igual a \"{{Value}}\"",
			"not_equivalent_to": "{{title}} no puede ser igual a \"{{Value}}\"",
			"equal_to":          "{{title}} debe ser igual a \"{{Value}}\"",
			"not_equal_to":      "{{title}} no puede ser igual a \"{{Value}}\"",
			"matching_to":       "{{title}} debe corresponder a \"{{Value}}\"",
			"not_matching_to":   "{{title}} no puede corresponder a \"{{Value}}\"",
			"identical_to":      "{{title}} debe ser igual a \"{{Value}}\"",
			"not_identical_to":  "{{title}} no puede ser igual a \"{{Value}}\"",
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
