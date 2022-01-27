package checker_test

import (
	"context"
	"database/sql"
)

func _(ctx context.Context, rows *sql.Rows) error {
	/*! assign rows.Err() to a variable and check it afterwards */
	if rows.Err() != nil {
		return rows.Err()
	}

	/*! assign ctx.Err() to a variable and check it afterwards */
	if ctx.Err() != nil {
		err := ctx.Err()
		panic(err)
	}

	switch {
	/*! assign ctx.Err() to a variable and check it afterwards */
	case ctx.Err() == nil:
	/*! assign rows.Err() to a variable and check it afterwards */
	case rows.Err() != nil:
	}

	/*! assign do() to a variable and check it afterwards */
	if do() == nil {
		panic("do nil")
	}

	/*! assign ctx.Err() to a variable and check it afterwards */
	for ctx.Err() == nil {
		// do some stuff
	}

	return nil
}

func do() error { return nil }
