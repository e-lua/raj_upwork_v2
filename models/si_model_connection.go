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
	HOSTNAME := "svc-d24221ca-add5-4726-807c-41c60cf39ecc-ddl.aws-virginia-2.svc.singlestore.com"
	PORT := "3306"
	USERNAME := "admin"
	PASSWORD := "d^<#|)ElJgab,$l=lBot"
	DATABASE := "TickerDatabase"

	once_pg.Do(func() {

		connection := USERNAME + ":" + PASSWORD + "@tcp(" + HOSTNAME + ":" + PORT + ")/" + DATABASE + "?parseTime=true"
		si_db, _ := sql.Open("mysql", connection)
		s_db = si_db
	})

	return s_db
}
