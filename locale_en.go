package valgo

func getLocaleEn() *Locale {
	return &Locale{
		ErrorKeyAfter:    "{{title}} must be after \"{{value}}\"",
		ErrorKeyNotAfter: "{{title}} can't be after \"{{value}}\"",

		ErrorKeyAfterOrEqualTo:    "{{title}} must be after or equal to \"{{value}}\"",
		ErrorKeyNotAfterOrEqualTo: "{{title}} can't be after or equal to \"{{value}}\"",

		ErrorKeyBefore:    "{{title}} must be before \"{{value}}\"",
		ErrorKeyNotBefore: "{{title}} can't be before \"{{value}}\"",

		ErrorKeyBeforeOrEqualTo:    "{{title}} must be before or equal to \"{{value}}\"",
		ErrorKeyNotBeforeOrEqualTo: "{{title}} can't be before or equal to \"{{value}}\"",

		ErrorKeyBetween:    "{{title}} must be between \"{{min}}\" and \"{{max}}\"",
		ErrorKeyNotBetween: "{{title}} can't be a value between \"{{min}}\" and \"{{max}}\"",

		ErrorKeyBlank:    "{{title}} must be blank",
		ErrorKeyNotBlank: "{{title}} can't be blank",

		ErrorKeyEmpty:    "{{title}} must be empty",
		ErrorKeyNotEmpty: "{{title}} can't be empty",

		ErrorKeyEqualTo:    "{{title}} must be equal to \"{{value}}\"",
		ErrorKeyNotEqualTo: "{{title}} can't be equal to \"{{value}}\"",

		ErrorKeyFalse:    "{{title}} must be false",
		ErrorKeyNotFalse: "{{title}} must not be false",

		ErrorKeyGreaterOrEqualTo:    "{{title}} must be greater than or equal to \"{{value}}\"",
		ErrorKeyNotGreaterOrEqualTo: "{{title}} can't be greater than or equal to \"{{value}}\"",

		ErrorKeyGreaterThan:    "{{title}} must be greater than \"{{value}}\"",
		ErrorKeyNotGreaterThan: "{{title}} can't be greater than \"{{value}}\"",

		ErrorKeyInSlice:    "{{title}} is not valid",
		ErrorKeyNotInSlice: "{{title}} is not valid",

		ErrorKeyLength:    "{{title}} must have a length equal to \"{{length}}\"",
		ErrorKeyNotLength: "{{title}} must not have a length equal to \"{{length}}\"",

		ErrorKeyLengthBetween:    "{{title}} must have a length between \"{{min}}\" and \"{{max}}\"",
		ErrorKeyNotLengthBetween: "{{title}} must not have a length between \"{{min}}\" and \"{{max}}\"",

		ErrorKeyLessOrEqualTo:    "{{title}} must be less than or equal to \"{{value}}\"",
		ErrorKeyNotLessOrEqualTo: "{{title}} must not be less than or equal to \"{{value}}\"",

		ErrorKeyLessThan:    "{{title}} must be less than \"{{value}}\"",
		ErrorKeyNotLessThan: "{{title}} can't be less than \"{{value}}\"",

		ErrorKeyMatchingTo:    "{{title}} must match to \"{{regexp}}\"",
		ErrorKeyNotMatchingTo: "{{title}} can't match to \"{{regexp}}\"",

		ErrorKeyMaxLength:    "{{title}} must not have a length longer than \"{{length}}\"",
		ErrorKeyNotMaxLength: "{{title}} must not have a length shorter than or equal to \"{{length}}\"",

		ErrorKeyMinLength:    "{{title}} must not have a length shorter than \"{{length}}\"",
		ErrorKeyNotMinLength: "{{title}} must not have a length longer than or equal to \"{{length}}\"",

		ErrorKeyNil:    "{{title}} must be nil",
		ErrorKeyNotNil: "{{title}} must not be nil",

		ErrorKeyPassing:    "{{title}} is not valid",
		ErrorKeyNotPassing: "{{title}} is not valid",

		ErrorKeyTrue:    "{{title}} must be true",
		ErrorKeyNotTrue: "{{title}} must not be true",

		ErrorKeyZero:    "{{title}} must be zero",
		ErrorKeyNotZero: "{{title}} must not be zero",

		ErrorKeyPositive:    "{{title}} must be positive",
		ErrorKeyNotPositive: "{{title}} must not be positive",

		ErrorKeyNegative:    "{{title}} must be negative",
		ErrorKeyNotNegative: "{{title}} must not be negative",

		ErrorKeyZeroOrNil:    "{{title}} must be zero or nil",
		ErrorKeyNotZeroOrNil: "{{title}} must not be zero or nil",

		ErrorKeyPositiveOrNil:    "{{title}} must be positive or nil",
		ErrorKeyNotPositiveOrNil: "{{title}} must not be positive or nil",

		ErrorKeyNegativeOrNil:    "{{title}} must be negative or nil",
		ErrorKeyNotNegativeOrNil: "{{title}} must not be negative or nil",

		ErrorKeyNaN:    "{{title}} must be NaN",
		ErrorKeyNotNaN: "{{title}} must not be NaN",

		ErrorKeyInfinite:    "{{title}} must be infinite",
		ErrorKeyNotInfinite: "{{title}} must not be infinite",

		ErrorKeyFinite:    "{{title}} must be finite",
		ErrorKeyNotFinite: "{{title}} must not be finite",
	}
}
