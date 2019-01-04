package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(192.168.43.227:13306)/video_server?charset=utf8")
	// dbConn, err = sql.Open("mysql", "root:123456@tcp(172.20.244.167:13306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}