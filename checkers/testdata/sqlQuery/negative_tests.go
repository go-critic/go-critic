package checker_test

import (
	"database/sql"
)

func foo() (int, int) { return 0, 0 }

func notQueryCall() {
	_, _ = foo()
}

func queryResultIsUsed(db *sql.DB, qe QueryExecer, mydb *myDatabase) {
	const queryString = "SELECT * FROM users"

	rows1, err := db.Query(queryString)
	_ = rows1

	rows2, err := qe.Query(queryString)
	_ = rows2

	rows3, err := mydb.Query(queryString)
	_ = rows3

	_ = err
}

func execIsUsed(db *sql.DB, qe QueryExecer, mydb *myDatabase) {
	const queryString = "UPDATE users SET name = 'gopher'"

	var err error

	_, err = db.Exec(queryString)
	_, err = qe.Exec(queryString)
	_, err = mydb.Exec(queryString)

	_ = err
}
