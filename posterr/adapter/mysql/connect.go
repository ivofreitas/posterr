package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strider-backend-test.com/config"
	"strider-backend-test.com/log"
	"sync"
)

var (
	once sync.Once
	conn *sql.DB
)

func GetConn() *sql.DB {
	once.Do(func() {
		mySQL := config.Get().MySQL
		logger := log.GetLogger()

		var err error
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", mySQL.Username, mySQL.Password, mySQL.Host, mySQL.Database)
		conn, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			logger.WithError(err).Fatal()
			return
		}

		if err = conn.Ping(); err != nil {
			logger.WithError(err).Fatal()
			return
		}

		conn.SetMaxIdleConns(mySQL.PoolConn)
		conn.SetConnMaxLifetime(mySQL.ConnLifetime)
	})

	return conn
}
