package checker_test

import (
	"os"
)

func _(a, b, c string) {
	/*! filepath.Join(a, b) should be preferred to the a + string(os.PathSeparator) + b */
	_ = a + string(os.PathSeparator) + b

	/*! filepath.Join(a, b) should be preferred to the a + string(os.PathSeparator) + b */
	_ = a + string(os.PathSeparator) + b + c

	/*! filepath.Join(c + a, b) should be preferred to the c + a + string(os.PathSeparator) + b */
	_ = c + a + string(os.PathSeparator) + b + c
}
