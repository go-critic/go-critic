package checker_test

import (
	"database/sql"
)

type myDatabase struct {
	*sql.DB
}

type Rows struct{}

type Row struct{}

type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Queryx(query string, args ...interface{}) (*Rows, error)
	QueryRowx(query string, args ...interface{}) *Row
}

type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type QueryExecer interface {
	Queryer
	Execer
}

type QueryExecerAlias = QueryExecer

func resultIgnored(db *sql.DB, q Queryer, qe QueryExecer, qea QueryExecerAlias, mydb *myDatabase) {
	const queryString = "UPDATE users SET name = 'gopher'"

	var err error

	/*! use db.Exec() if returned result is not needed */
	_, err = db.Query(queryString)

	/*! use qe.Exec() if returned result is not needed */
	_, err = qe.Query(queryString)

	/*! use qe.Exec() if returned result is not needed */
	_, err = qe.Queryx(queryString)

	/*! use mydb.Exec() if returned result is not needed */
	_, err = mydb.Query(queryString)

	/*! ignoring Query() rows result may lead to a connection leak */
	_, err = q.Query(queryString)

	/*! use qea.Exec() if returned result is not needed */
	_, err = qea.Query(queryString)

	/*! ignoring Query() rows result may lead to a connection leak */
	_, err = q.Queryx(queryString)

	_ = err
}
