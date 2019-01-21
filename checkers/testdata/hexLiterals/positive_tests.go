package checker_test

func someTest() {
	/*! Should be 0x12 */
	_ = 0X12

	/*! Should be 0xff or 0xFF */
	_ = 0xfF

	/*! Should be 0xEE */
	_ = 0XEE

	/*! Should be 0xaa */
	_ = 0Xaa
}
