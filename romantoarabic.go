func arabicToRoman(arabic int) (string, error) {
	if arabic <= 0 || arabic > 100 {
		return "", fmt.Errorf("Римские числа не могут быть меньше 1")
	}

	var romanNumerals = []struct {
		Value  int
		Symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""

	for _, numeral := range romanNumerals {
		for arabic >= numeral.Value {
			roman += numeral.Symbol
			arabic -= numeral.Value
		}
	}

	return roman, nil
}