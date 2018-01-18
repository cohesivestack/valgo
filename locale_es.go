package valgo

func init() {
	getLocales()["es"] = locale{
		Messages: map[string]string{
			"blank":     "\"{{Title}}\" debe estar en blanco",
			"not_blank": "\"{{Title}}\" no puede estar en blanco",
			"empty":     "\"{{Title}}\" debe estar vacío",
			"not_empty": "\"{{Title}}\" no puede estar vacío",
		},
	}
}
