package mysql

import (
	"database/sql"

	"github.com/pkg/errors"
)

var (
	mysqlConnectionString = "root:123456@tcp(127.0.0.1:3306)/test"
)

type TableResult struct {
	Id   int64
	Name string
}

func Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", mysqlConnectionString)
	if err != nil {
		return nil, errors.Wrap(err, "open database fail")
	}
	return db, nil
}

func Query(sqlString string, db *sql.DB) ([]TableResult, error) {
	rows, err := db.Query(sqlString, 1)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "query database fail")
	}
	var tableResultSlice []TableResult

	for rows.Next() {
		var tableResult TableResult
		err := rows.Scan(&tableResult.Id, &tableResult.Name)
		if err != nil {
			return nil, errors.Wrap(err, "scan row fail")
		}

		tableResultSlice = append(tableResultSlice, tableResult)
	}

	return tableResultSlice, nil
}
