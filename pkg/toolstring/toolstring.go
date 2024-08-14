package toolstring

/*
"123456789",5=>"123.."
"1234",5=>"1234"
"123456",5=>"123.."
"12",5=>"12"

*/
func Minimize(str string, maxLen int) string {
	strLen := len(str)
	if strLen > maxLen && strLen > 2 {
		return str[:maxLen-2] + ".."
	}
	return str
}

func Left(str string, n int) string {
	if n >= len(str) {
		return str
	}
	return str[:n]
}

func Right(str string, n int) string {
	if n >= len(str) {
		return str
	}
	return str[len(str)-n:]
}
