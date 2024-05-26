package mysqldb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "msql"
)

type Mysql struct {
	db *sql.DB
}

func New(dbUser, dbPassword, dbHost, dbPort, dbName string) (*Mysql, error) {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatalf("mysqldb connect failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("mysqldb ping failure: %v", err)
	}

	return &Mysql{db: db}, nil
}

func (this Mysql) Close() {
	err := this.db.Close()
	if err != nil {
		log.Fatalf("mysqldb close failure: %v", err)
	}
}

func (this Mysql) InsertUser(userName string) error {
	this.db.Exec("INSERT...")

	return nil
}

func (this Mysql) SelectSingleUser(userName string) (string, error) {
	this.db.Exec("SELECT...")

	return "user", nil
}