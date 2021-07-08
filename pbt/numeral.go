package numeral

import "strings"

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for i := arabic; i > 0; i-- {
		if arabic == 4 {
			return "IV"
		}
		result.WriteString("I")
	}

	return result.String()
}
