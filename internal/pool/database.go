package pool

import (
	"database/sql"
	"log"

	"github.com/golang/glog"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Database *sql.DB
)

func init() {
	db, err := sql.Open("sqlite3", "./aaa.db")
	if err != nil {
		glog.Errorf("connect db fail: %s\n", err)
		log.Panic(err.Error())
	}
	Database = db
}
