package sqlite

import "database/sql"

func openDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", "../../../../chinook.sqlite")
}
