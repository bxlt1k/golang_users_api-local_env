package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"time"
)

func CreateConnection(mysqlConnStr string) (*sqlx.DB, error) {
	var mysqlConn *sqlx.DB
	var err error

	for i := 0; i < 10; i++ {
		log.Println("Try to connect to mysql... Attempt " + strconv.Itoa(i+1))
		mysqlConn, err = sqlx.Connect("mysql", mysqlConnStr)
		if err == nil {
			log.Println("Connection with mysql created")
			return mysqlConn, nil
		}
		time.Sleep(5 * time.Second)
	}

	return nil, err
}
