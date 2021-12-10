package checker_test

import "time"

func _(t time.Time, tp *time.Time) {
	/*! use t.UnixMilli() instead of t.Unix() / 1000 */
	_ = t.Unix() / 1000
	/*! use tp.UnixMilli() instead of tp.Unix() / 1000 */
	_ = tp.Unix() / 1000

	/*! use t.UnixMicro() instead of t.UnixNano() * 1000 */
	_ = t.UnixNano() * 1000
	/*! use tp.UnixMicro() instead of tp.UnixNano() * 1000 */
	_ = tp.UnixNano() * 1000
}
