package checker_test

import (
	"context"
	"database/sql"
)

func _(ctx context.Context, rows *sql.Rows) error {
	if err := rows.Err(); err != nil {
		return rows.Err()
	}

	if err2 := ctx.Err(); err2 != nil {
		panic(err2)
	}

	isErr1 := ctx.Err() == nil
	isErr2 := rows.Err() != nil

	switch {
	case isErr1:
	case isErr2:
	}

	if do() == context.Canceled {
		panic("do nil")
	}

	for err := ctx.Err(); err == nil; err = ctx.Err() {
		// do some stuff
	}

	return nil
}
