package checker_tests

import "database/sql"

func normalLocalUseWithDefer() {
    db, err := sql.Open("postgres", "")
    if err != nil {
        return
    }

    rows, err := db.Query("SELECT * FROM testtable")
    if err != nil {
        return
    }
    defer rows.Close()

    for rows.Next() {
        var testdata string
        rows.Scan(&testdata)
    }

    return
}

/*
func normalLocalUseWithAnonDefer() {
    db, err := sql.Open("postgres", "")
    if err != nil {
        return
    }

    rows, err := db.Query("SELECT * FROM testtable")
    if err != nil {
        return
    }
    defer func() {
        rows.Close()
    }()

    for rows.Next() {
        var testdata string
        rows.Scan(&testdata)
    }

    return
}
*/

func normalLocalUseWithoutDefer() {
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
    rows.Close()

    return
}

func normalUseWhereRowsFromOtherMethodWithDefer() {
    db, err := sql.Open("postgres", "")
    if err != nil {
        return
    }

    rows, err := testMethodReturningRows(db)
    if err != nil {
        return
    }
    defer rows.Close()

    for rows.Next() {
        var testdata string
        rows.Scan(&testdata)
    }

    return
}

func normalUseWhereRowsFromOtherMethodWithoutDefer() {
    db, err := sql.Open("postgres", "")
    if err != nil {
        return
    }

    rows, err := testMethodReturningRows(db)
    if err != nil {
        return
    }

    for rows.Next() {
        var testdata string
        rows.Scan(&testdata)
    }
    rows.Close()

    return
}


func normalUseWhereRowsFromOtherMethodToOtherMethodWithDefer() {
    db, err := sql.Open("postgres", "")
    if err != nil {
        return
    }

    rows, err := testMethodReturningRows(db)
    if err != nil {
        return
    }

    testMethodCloseRowsWithDefer(rows)

    return
}

func normalUseWhereRowsFromOtherMethodToOtherMethodWithoutDefer() {
    db, err := sql.Open("postgres", "")
    if err != nil {
        return
    }

    rows, err := testMethodReturningRows(db)
    if err != nil {
        return
    }

    testMethodCloseRowsWithoutDefer(rows)

    return
}

// Internal methods

func testMethodReturningRows(db *sql.DB) (*sql.Rows, error) {
    rows, err := db.Query("SELECT * FROM testtable")
    if err != nil {
        return nil, err
    }

    return rows, nil
}

func testMethodCloseRowsWithDefer(rows *sql.Rows) {
    defer rows.Close()

    for rows.Next() {
        var testdata string
        rows.Scan(&testdata)
    }
}

func testMethodCloseRowsWithoutDefer(rows *sql.Rows) {
    for rows.Next() {
        var testdata string
        rows.Scan(&testdata)
    }
    rows.Close()
}

func testPosMethodReturningRows(db *sql.DB) (*sql.Rows, error) {
    rows, err := db.Query("SELECT * FROM testtable")
    if err != nil {
        return nil, err
    }

    return rows, nil
}
