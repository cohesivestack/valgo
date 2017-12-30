package valgo

func init() {
	getLocales()["en"] = locale{
		Messages: map[string]string{
			"blank":     "{{Title}} should be blank",
			"not_blank": "{{Title}} should not be blank",
			"empty":     "{{Title}} should be empty",
			"not_empty": "{{Title}} should not be empty",
		},
	}
}
