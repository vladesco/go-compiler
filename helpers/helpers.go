package helpers

func IsLetter(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || char == '_'
}

func IsWhiteSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func IsDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
