package valgo

func getLocaleHu() *Locale {
	return &Locale{
		ErrorKeyAfter:    "{{title}} csak \"{{value}}\" után következhet",
		ErrorKeyNotAfter: "{{title}} nem következhet \"{{value}}\" után",

		ErrorKeyAfterOrEqualTo:    "{{title}} meg kell egyezzen vagy követnie kell \"{{value}}\" értékét",
		ErrorKeyNotAfterOrEqualTo: "{{title}} meg kell egyezzen vagy meg kell előzze \"{{value}}\" értékét",

		ErrorKeyBefore:    "{{title}} meg kell előzze \"{{value}}\" értékét",
		ErrorKeyNotBefore: "{{title}} nem előzheti meg \"{{value}}\" értékét",

		ErrorKeyBeforeOrEqualTo:    "{{title}} meg kell egyezzen vagy meg kell előzze \"{{value}}\" értékét",
		ErrorKeyNotBeforeOrEqualTo: "{{title}} meg kell egyezzen vagy nem előzheti meg \"{{value}}\" értékét",

		ErrorKeyBetween:    "{{title}} \"{{min}}\" és \"{{max}}\" közé kell essen",
		ErrorKeyNotBetween: "{{title}} nem eshet \"{{min}}\" és \"{{max}}\" közé",

		ErrorKeyBlank:    "{{title}} csak szóközökből állhat",
		ErrorKeyNotBlank: "{{title}} nem állhat csak szóközökből",

		ErrorKeyEmpty:    "{{title}} üres kell legyen",
		ErrorKeyNotEmpty: "{{title}} nem lehet üres",

		ErrorKeyEqualTo:    "{{title}} meg kell egyezzen \"{{value}}\" értékével",
		ErrorKeyNotEqualTo: "{{title}} nem egyezhet meg \"{{value}}\" értékével",

		ErrorKeyFalse:    "{{title}} hamis kell legyen",
		ErrorKeyNotFalse: "{{title}} nem lehet hamis",

		ErrorKeyGreaterOrEqualTo:    "{{title}} nagyobb vagy egyenlő \"{{value}}\" értékénél",
		ErrorKeyNotGreaterOrEqualTo: "{{title}} nem lehet nagyobb vagy egyenlő \"{{value}}\" értékénél",

		ErrorKeyGreaterThan:    "{{title}} nagyobb kell legyen \"{{value}}\" értékénél",
		ErrorKeyNotGreaterThan: "{{title}} nem lehet nagyobb \"{{value}}\" értékénél",

		ErrorKeyInSlice:    "{{title}} nincs az elfogadható értékek között",
		ErrorKeyNotInSlice: "{{title}} nem lehet a kizárt értékek között",

		ErrorKeyLength:    "{{title}} hossza egyenlő kell legyen \"{{length}}\" értékével",
		ErrorKeyNotLength: "{{title}} nem lehet egyenlő \"{{length}}\" értékével",

		ErrorKeyLengthBetween:    "{{title}} hossza \"{{min}}\" és \"{{max}}\" közé kell essen",
		ErrorKeyNotLengthBetween: "{{title}} hossza nem eshet \"{{min}}\" és \"{{max}}\" közé",

		ErrorKeyLessOrEqualTo:    "{{title}} kevesebb vagy egyenlő \"{{value}}\" értéknél",
		ErrorKeyNotLessOrEqualTo: "{{title}} nem lehet kevesebb vagy egyenlő \"{{value}}\" értéknél",

		ErrorKeyLessThan:    "{{title}} kevesebb kell legyen, mint \"{{value}}\"",
		ErrorKeyNotLessThan: "{{title}} nem lehet kevesebb, mint \"{{value}}\"",

		ErrorKeyMatchingTo:    "{{title}} meg kell feleljen a \"{{regexp}}\" mintának",
		ErrorKeyNotMatchingTo: "{{title}} nem felelhet meg a \"{{regexp}}\" mintának",

		ErrorKeyMaxLength:    "{{title}} nem lehet hosszabb \"{{length}}\" értékénél",
		ErrorKeyNotMaxLength: "{{title}} nem lehet hosszabb vagy egyenlő \"{{length}}\" értékével",

		ErrorKeyMinLength:    "{{title}} nem lehet rövidebb \"{{length}}\" értékénél",
		ErrorKeyNotMinLength: "{{title}} nem lehet rövidebb vagy egyenlő \"{{length}}\" értékénél",

		ErrorKeyNil:    "{{title}} nil kell legyen",
		ErrorKeyNotNil: "{{title}} nem lehet nil",

		ErrorKeyPassing:    "{{title}} nem érvényes",
		ErrorKeyNotPassing: "{{title}} nem érvényes",

		ErrorKeyTrue:    "{{title}} igaz kell legyen",
		ErrorKeyNotTrue: "{{title}} nem lehet igaz",

		ErrorKeyZero:    "{{title}} nulla kell legyen",
		ErrorKeyNotZero: "{{title}} nem lehet nulla",

		ErrorKeyPositive:    "{{title}} pozitív kell legyen",
		ErrorKeyNotPositive: "{{title}} nem lehet pozitív",

		ErrorKeyNegative:    "{{title}} negatív kell legyen",
		ErrorKeyNotNegative: "{{title}} nem lehet negatív",

		ErrorKeyZeroOrNil:    "{{title}} nulla vagy nil kell legyen",
		ErrorKeyNotZeroOrNil: "{{title}} nem lehet nulla vagy nil",

		ErrorKeyPositiveOrNil:    "{{title}} pozitív vagy nil kell legyen",
		ErrorKeyNotPositiveOrNil: "{{title}} nem lehet pozitív vagy nil",

		ErrorKeyNegativeOrNil:    "{{title}} negatív vagy nil kell legyen",
		ErrorKeyNotNegativeOrNil: "{{title}} nem lehet negatív vagy nil",

		ErrorKeyNaN:    "{{title}} NaN kell legyen",
		ErrorKeyNotNaN: "{{title}} nem lehet NaN",

		ErrorKeyInfinite:    "{{title}} végtelen kell legyen",
		ErrorKeyNotInfinite: "{{title}} nem lehet végtelen",

		ErrorKeyFinite:    "{{title}} véges kell legyen",
		ErrorKeyNotFinite: "{{title}} nem lehet véges",
	}
}
