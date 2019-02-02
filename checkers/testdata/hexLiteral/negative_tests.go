package checker_test

func negatives() {
	_ = 0xFF
	_ = 0xff
	_ = "0xfF"
}

func noLetters() {
	_ = 0x11
	_ = 0x0
}

func digitsAndLetters() {
	_ = 0x1aa
	_ = 0x1AA
}

func decimals() {
	_ = 1
	_ = 10
	_ = 12345
}
