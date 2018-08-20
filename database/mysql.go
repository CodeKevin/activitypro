package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Orm *xorm.Engine

func init() {
	var err error
	Orm, err = xorm.NewEngine("mysql", "root:root@tcp(localhost:32768)/activitypro?charset=utf8")
	Orm.ShowSQL(true)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var SqlDB *sql.DB

func init() {
	var err error
	SqlDB, err = sql.Open("mysql", "root:root@tcp(localhost:32768)/activitypro?charset=utf8")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

}
