package valgo

func getLocaleEs() *Locale {
	return &Locale{
		ErrorKeyBetween:    "{{title}} debe estar entre \"{{min}}\" y \"{{max}}\"",
		ErrorKeyNotBetween: "{{title}} no puede ser un valor entre \"{{min}}\" y \"{{max}}\"",

		ErrorKeyBlank:    "{{title}} debe estar en blanco",
		ErrorKeyNotBlank: "{{title}} no puede estar en blanco",

		ErrorKeyEmpty:    "{{title}} debe estar vacío",
		ErrorKeyNotEmpty: "{{title}} no puede estar vacío",

		ErrorKeyEqualTo:    "{{title}} debe ser igual a \"{{value}}\"",
		ErrorKeyNotEqualTo: "{{title}} no puede ser igual a \"{{value}}\"",

		ErrorKeyFalse:    "{{title}} debe ser falso",
		ErrorKeyNotFalse: "{{title}} no debe ser falso",

		ErrorKeyGreaterOrEqualTo:    "{{title}} debe ser mayor o igual a \"{{value}}\"",
		ErrorKeyNotGreaterOrEqualTo: "{{title}} no puede ser mayor o igual a \"{{value}}\"",

		ErrorKeyGreaterThan:    "{{title}} debe ser mayor que \"{{value}}\"",
		ErrorKeyNotGreaterThan: "{{title}} no puede ser mayor que \"{{value}}\"",

		ErrorKeyInSlice:    "{{title}} no es válido",
		ErrorKeyNotInSlice: "{{title}} no es válido",

		ErrorKeyLength:    "{{title}} debe tener una longitud igual a \"{{length}}\"",
		ErrorKeyNotLength: "{{title}} no debe tener una longitud igual a \"{{length}}\"",

		ErrorKeyLengthBetween:    "{{title}} debe tener una longitud entre \"{{min}}\" and \"{{max}}\"",
		ErrorKeyNotLengthBetween: "{{title}} no debe tener una longitud entre \"{{min}}\" and \"{{max}}\"",

		ErrorKeyLessOrEqualTo:    "{{title}} debe ser menor o igual a \"{{value}}\"",
		ErrorKeyNotLessOrEqualTo: "{{title}} no debe ser menor o igual a \"{{value}}\"",

		ErrorKeyLessThan:    "{{title}} debe ser menor que \"{{value}}\"",
		ErrorKeyNotLessThan: "{{title}} no puede ser menor que \"{{value}}\"",

		ErrorKeyMatchingTo:    "{{title}} debe coincidir con \"{{regexp}}\"",
		ErrorKeyNotMatchingTo: "{{title}} no puede coincidir con \"{{regexp}}\"",

		ErrorKeyMaxLength:    "{{title}} no debe tener una longitud mayor a \"{{length}}\"",
		ErrorKeyNotMaxLength: "{{title}} no debe tener una longitud menor o igual a \"{{length}}\"",

		ErrorKeyMinLength:    "{{title}} no debe tener una longitud menor a \"{{length}}\"",
		ErrorKeyNotMinLength: "{{title}} no debe tener una longitud mayor o igual a \"{{length}}\"",

		ErrorKeyNil:    "{{title}} debe ser nulo",
		ErrorKeyNotNil: "{{title}} no debe ser nulo",

		ErrorKeyPassing:    "{{title}} no es válido",
		ErrorKeyNotPassing: "{{title}} no es válido",

		ErrorKeyTrue:    "{{title}} debe ser verdadero",
		ErrorKeyNotTrue: "{{title}} no debe ser verdadero",

		ErrorKeyZero:    "{{title}} debe ser cero",
		ErrorKeyNotZero: "{{title}} no debe ser cero",
	}
}
