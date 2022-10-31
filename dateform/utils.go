package dateform

func toDigit(str string, length int) string {
	result := str

	for len(result) < length {
		result = "0" + result
	}

	return result
}
