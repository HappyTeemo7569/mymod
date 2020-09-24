package base

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func getDb() *sqlx.DB {
	return Db
}

func InitMysqlNormal() {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		configBase.MySql.Auth,
		configBase.MySql.Pwd,
		configBase.MySql.Addr,
		configBase.MySql.Port,
		configBase.MySql.Db)

	var err error
	Db, err = sqlx.Open("mysql", connstr)
	if err != nil {
		panic(err)
		//Logger.Errorf("open mysql failed,", err)
		return
	}
}
