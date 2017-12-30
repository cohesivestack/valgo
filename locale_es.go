package valgo

func init() {
	getLocales()["es"] = locale{
		Messages: map[string]string{
			"blank":     "{{Title}} debería estar en blanco",
			"not_blank": "{{Title}} no debería estar en blanco",
			"empty":     "{{Title}} debería estar vacío",
			"not_empty": "{{Title}} no debería estar vacío",
		},
	}
}
