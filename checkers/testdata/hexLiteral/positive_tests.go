package checker_test

func bad0X() {
	/*! prefer 0x over 0X, s/0X12/0x12/ */
	_ = 0X12
	/*! prefer 0x over 0X, s/0XEE/0xEE/ */
	_ = 0XEE
	/*! prefer 0x over 0X, s/0Xaa/0xaa/ */
	_ = 0Xaa
}

func mixedLetterDigits() {
	/*! don't mix hex literal letter digits casing */
	_ = 0xfF
	/*! don't mix hex literal letter digits casing */
	_ = 0xFf
	/*! don't mix hex literal letter digits casing */
	_ = 0x11f0F
	/*! don't mix hex literal letter digits casing */
	_ = 0xff11FF
	/*! don't mix hex literal letter digits casing */
	_ = 0xabcdE
}
