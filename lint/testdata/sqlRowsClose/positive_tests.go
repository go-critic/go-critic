package checker_tests

import "database/sql"

/// local variable db.Rows have not Close call
func normalPosLocalUse() {
	db, err := sql.Open("postgres", "")
	if err != nil {
		return
	}

	rows, err := db.Query("SELECT * FROM testtable")
	if err != nil {
		return
	}

	for rows.Next() {
		var testdata string
		rows.Scan(&testdata)
	}

	return
}

/// local variable db.Rows have not Close call
func normalPosLocalUseWithCall() {
	db, err := sql.Open("postgres", "")
	if err != nil {
		return
	}

	rows, err := db.Query("SELECT * FROM testtable")
	if err != nil {
		return
	}

	for rows.Next() {
		testPosMethodCallRows(rows)
	}

	return
}

func testPosMethodCallRows(rows *sql.Rows) {
	for rows.Next() {
		var testdata string
		rows.Scan(&testdata)
	}
}
