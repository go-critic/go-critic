package timeCmpSimplify

import (
	"time"
)

func negative(x, y time.Time) {
	if x.Before(y) {
		print(1)
	}

	if y.After(x) {
		print(1)
	}

	print(x.After(y))
}
