package valgo

func getLocaleDe() *Locale {
	return &Locale{
		ErrorKeyAfter:    "{{title}} muss nach \"{{value}}\" sein",
		ErrorKeyNotAfter: "{{title}} darf nicht nach \"{{value}}\" sein",

		ErrorKeyAfterOrEqualTo:    "{{title}} muss nach oder gleich \"{{value}}\" sein",
		ErrorKeyNotAfterOrEqualTo: "{{title}} darf nicht nach oder gleich \"{{value}}\" sein",

		ErrorKeyBefore:    "{{title}} muss vor \"{{value}}\" sein",
		ErrorKeyNotBefore: "{{title}} darf nicht vor \"{{value}}\" sein",

		ErrorKeyBeforeOrEqualTo:    "{{title}} muss vor oder gleich \"{{value}}\" sein",
		ErrorKeyNotBeforeOrEqualTo: "{{title}} darf nicht vor oder gleich \"{{value}}\" sein",

		ErrorKeyBetween:    "{{title}} muss zwischen \"{{min}}\" und \"{{max}}\" sein",
		ErrorKeyNotBetween: "{{title}} darf nicht zwischen \"{{min}}\" und \"{{max}}\" sein",

		ErrorKeyBlank:    "{{title}} darf nicht ausgefüllt sein",
		ErrorKeyNotBlank: "{{title}} muss ausgefüllt sein",

		ErrorKeyEmpty:    "{{title}} muss leer sein",
		ErrorKeyNotEmpty: "{{title}} darf nicht leer sein",

		ErrorKeyEqualTo:    "{{title}} muss identisch zu \"{{value}}\" sein",
		ErrorKeyNotEqualTo: "{{title}} darf nicht identisch zu \"{{value}}\" sein",

		ErrorKeyFalse:    "{{title}} muss \"false\" sein",
		ErrorKeyNotFalse: "{{title}} darf nicht \"false\" sein",

		ErrorKeyGreaterOrEqualTo:    "{{title}} muss größer oder gleich als \"{{value}}\" sein",
		ErrorKeyNotGreaterOrEqualTo: "{{title}} darf nicht größer oder gleich als \"{{value}}\" sein",

		ErrorKeyGreaterThan:    "{{title}} muss größer als \"{{value}}\" sein",
		ErrorKeyNotGreaterThan: "{{title}} darf nicht größer als \"{{value}}\" sein",

		ErrorKeyInSlice:    "{{title}} ist nicht gültig",
		ErrorKeyNotInSlice: "{{title}} ist nicht gültig",

		ErrorKeyLength:    "{{title}} muss exakt \"{{length}}\" Zeichen lang sein",
		ErrorKeyNotLength: "{{title}} darf nicht \"{{length}}\" Zeichen lang sein",

		ErrorKeyLengthBetween:    "{{title}} muss zwischen \"{{min}}\" und \"{{max}}\" Zeichen lang sein",
		ErrorKeyNotLengthBetween: "{{title}} darf nicht zwischen \"{{min}}\" und \"{{max}}\" Zeichen lang sein",

		ErrorKeyLessOrEqualTo:    "{{title}} muss kleiner oder gleich als \"{{value}}\" sein",
		ErrorKeyNotLessOrEqualTo: "{{title}} darf nicht kleiner oder gleich als \"{{value}}\" sein",

		ErrorKeyLessThan:    "{{title}} muss weniger als \"{{value}}\" sein",
		ErrorKeyNotLessThan: "{{title}} darf nicht weniger als \"{{value}}\" sein",

		ErrorKeyMatchingTo:    "{{title}} muss \"{{regexp}}\" entsprechen",
		ErrorKeyNotMatchingTo: "{{title}} darf nicht \"{{regexp}}\" entsprechen",

		ErrorKeyMaxLength:    "{{title}} darf nicht länger als \"{{length}}\" sein",
		ErrorKeyNotMaxLength: "{{title}} muss länger als \"{{length}}\" sein",

		ErrorKeyMinLength:    "{{title}} darf nicht kürzer als \"{{length}}\" sein",
		ErrorKeyNotMinLength: "{{title}} muss kürzer als \"{{length}}\" sein",

		ErrorKeyNil:    "{{title}} muss \"nil\" sein",
		ErrorKeyNotNil: "{{title}} darf nicht \"nil\" sein",

		ErrorKeyPassing:    "{{title}} ist nicht gültig",
		ErrorKeyNotPassing: "{{title}} ist nicht gültig",

		ErrorKeyTrue:    "{{title}} muss \"true\" sein",
		ErrorKeyNotTrue: "{{title}} darf nicht \"true\" sein",

		ErrorKeyZero:    "{{title}} muss 0 sein",
		ErrorKeyNotZero: "{{title}} darf nicht 0 sein",
	}
}
