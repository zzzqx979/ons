package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func DatabaseCon() {
	db, err := xorm.NewEngine(viper.GetString("db.type"), viper.GetString("db.source"))
	if err != nil {
		panic(fmt.Errorf("fatal errors connect database: %w", err))
	}
	db.SetLogLevel(0)
	db.ShowSQL(true)
	engine = db
}

func SyncAIotTables() {
	err := GetDB().Sync2(&OnsInfo{}, &ObjectModel{}, &Factory{}, &Product{})
	if err != nil {
		panic(fmt.Errorf("fatal errors sync aiot tables: %w", err))
	}
}

func SyncGatewayTables() {
	err := GetDB().Sync2(&Upstream{})
	if err != nil {
		panic(fmt.Errorf("fatal errors sync gateway tables: %w", err))
	}
}

func GetDB() *xorm.Engine {
	return engine
}
