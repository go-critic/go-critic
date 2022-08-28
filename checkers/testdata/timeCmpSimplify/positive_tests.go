package timeCmpSimplify

import (
    "time"
)

func _(x, y time.Time) {
    /*! suggestion: x.After(y) */
    if !x.Before(y) {
        print(42)
    }

    /*! suggestion: y.Before(x) */
    if !y.After(x) {
        print(51)
    }

    /*! suggestion: y.After(x) */
    print(!y.Before(x))
}
