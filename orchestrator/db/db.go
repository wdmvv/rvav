package db

// for managing database connection & requests

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type DBData struct{
	user string
	password string
	dbname string
	tabName string
	conn *sql.DB
}

var DBCOnn *DBData

func GetConn(user, password, dbname, tabname string) error{
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		return fmt.Errorf("failed to connect to database")
	}

	var dbdata DBData
	dbdata.user = user
	dbdata.password = password
	dbdata.dbname = dbname
	dbdata.tabName = tabname
	dbdata.conn = db

	DBCOnn = &dbdata
	return nil
}

func (db *DBData) Table() error {
    query := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            expr VARCHAR(255) PRIMARY KEY,
            result REAL
        )`, db.tabName)

    _, err := db.conn.Exec(query)
    if err != nil {
        return err
    }
    return nil
}

func (db *DBData) SetExpr(expr string, val float64) error {
    query := fmt.Sprintf("INSERT INTO %s (expr, result) VALUES ($1, $2)", db.tabName)
    _, err := db.conn.Exec(query, expr, val)
    if err != nil {
        return err
    }
    return nil
}

func (db *DBData) GetExpr(expr string) (float64, error) {
    var result float64
    query := fmt.Sprintf("SELECT result FROM %s WHERE expr = $1", db.tabName)
    row := db.conn.QueryRow(query, expr)
    err := row.Scan(&result)
    if err != nil {
        return 0, err
    }
    return result, nil
}

func (db *DBData) ExprExists(expr string) (bool, error) {
    query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE expr = $1)", db.tabName)
    var exists bool
    err := db.conn.QueryRow(query, expr).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}