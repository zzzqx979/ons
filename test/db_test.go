package test

import (
	"fmt"
	"github.com/spf13/viper"
	"ons/db"
	"testing"
)

func dbInit() {
	// 设置配置文件
	viper.SetConfigFile("../app.yaml")
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal config file errors: %w", err))
	}
	db.DatabaseCon()
	//db.SyncAIotTables()
}

func TestGetObjectModelByCodes(t *testing.T) {
	dbInit()
	t.Log(db.GetObjectModelByCodes("321", "123"))
}
