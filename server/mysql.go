package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func databaseError(err error) {
	if err != nil {
		serverLogger("Database error", err.Error(), "Error")
	}
}

func checkToken(cliToken string) string {
	// Make connect string
	dbString := dbConf.username + ":" + dbConf.password + "@tcp(" + dbConf.addr + ":" + dbConf.port + ")/LAM?charset=utf8&parseTime=true"
	// Connect
	db, err := sql.Open("mysql", dbString)
	databaseError(err)

	// Query
	res, err := db.Query("SELECT * FROM token WHERE token=?")
	databaseError(err)

	db.Close()

	// Extract data
	for res.Next() {
		var tag string
		var token string
		var timestamp string
		err := res.Scan(&tag, &token, &timestamp)
		databaseError(err)
		// Compare
		if token == cliToken {
			return tag
		}
	}

	// Invalid token
	return ""
}

func checkTagDuplicate(cliTag string) bool {
	// Make connect string
	dbString := dbConf.username + ":" + dbConf.password + "@tcp(" + dbConf.addr + ":" + dbConf.port + ")/LAM?charset=utf8&parseTime=true"
	// Connect
	db, err := sql.Open("mysql", dbString)
	databaseError(err)

	// Query
	res, err := db.Query("SELECT * FROM token")
	databaseError(err)

	db.Close()

	// Extract data
	for res.Next() {
		var tag string
		var token string
		var timestamp string
		err := res.Scan(&tag, &token, &timestamp)
		databaseError(err)
		// Compare
		if tag == cliTag {
			return true
		}
	}

	// No duplicate tag
	return false
}

func storeMessage(tag string, content string) bool {
	// Connect
	db, err := sql.Open("mysql", "LAM:lamLAM@tcp(127.0.0.1:3306)/LAM?charset=utf8&parseTime=true")
	databaseError(err)

	// Statement
	stmt, err := db.Prepare("INSERT message SET tag=?, content=?, timestamp=?, ifRead=?")
	databaseError(err)

	// Insert
	res, err := stmt.Exec(tag, content, time.Now().Format("2006-01-02 15:04:05"), false)
	databaseError(err)

	db.Close()

	// Validate
	num, err := res.RowsAffected()
	databaseError(err)

	if num == 1 {
		return true
	}
	return false
}

func storeToken(tag string, token string) bool {
	// Connect
	db, err := sql.Open("mysql", "LAM:lamLAM@tcp(127.0.0.1:3306)/LAM?charset=utf8&parseTime=true")
	databaseError(err)

	// Statement
	stmt, err := db.Prepare("INSERT token SET tag=?, token=?, timestamp=?")
	databaseError(err)

	// Insert
	res, err := stmt.Exec(tag, token, time.Now().Format("2006-01-02 15:04:05"))
	databaseError(err)

	db.Close()

	// Validate
	num, err := res.RowsAffected()
	databaseError(err)

	if num == 1 {
		return true
	}
	return false
}

func delToken(cliTag string) bool {
	// Connect
	db, err := sql.Open("mysql", "LAM:lamLAM@tcp(127.0.0.1:3306)/LAM?charset=utf8&parseTime=true")
	databaseError(err)

	// Statement
	stmt, err := db.Prepare("DELETE FROM token WHERE tag=?")
	databaseError(err)

	// Insert
	res, err := stmt.Exec(cliTag)
	databaseError(err)

	db.Close()

	// Validate
	num, err := res.RowsAffected()
	databaseError(err)

	if num == 1 {
		return true
	}
	return false
}

func fetchToken() []string {
	// Connect
	db, err := sql.Open("mysql", "LAM:lamLAM@tcp(127.0.0.1:3306)/LAM?charset=utf8&parseTime=true")
	databaseError(err)

	// Query
	res, err := db.Query("SELECT * FROM token")
	databaseError(err)

	db.Close()

	// Data array
	var tagNtoken []string
	// Extract data
	for res.Next() {
		var tag string
		var token string
		var timestamp string
		err := res.Scan(&tag, &token, &timestamp)
		databaseError(err)
		// Append data
		str := tag + ":" + token
		tagNtoken = append(tagNtoken, str)
	}
	return tagNtoken
}
