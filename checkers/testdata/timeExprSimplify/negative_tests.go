package checker_test

import "time"

func _(t time.Time, tp *time.Time) {
	_ = t.UnixMilli() / 1000
	_ = t.UnixMilli() * 1000
	_ = 1000 * t.UnixMilli()
	_ = t.UnixMilli()

	_ = t.UnixMilli() / 1000
	_ = t.UnixMilli() * 1000
	_ = 1000 * t.UnixMilli()
	_ = t.UnixMilli()

	_ = t.UnixMicro() * 1000
}
