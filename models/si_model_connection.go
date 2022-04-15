package models

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var SingleStoreCN = Conectar_Si_DB()

var (
	once_pg sync.Once
	s_db    *sql.DB
)

func Conectar_Si_DB() *sql.DB {

	// TODO: read from config
	HOSTNAME := "svc-67e316db-38ba-4a61-821b-cb5e53b4324c-ddl.aws-oregon-2.svc.singlestore.com"
	PORT := "3306"
	USERNAME := "admin"
	PASSWORD := "wdfsdf3123ga..FF"
	DATABASE := "TickerDatabase"

	once_pg.Do(func() {

		connection := USERNAME + ":" + PASSWORD + "@tcp(" + HOSTNAME + ":" + PORT + ")/" + DATABASE + "?parseTime=true"
		si_db, _ := sql.Open("mysql", connection)
		s_db = si_db
	})

	return s_db
}
