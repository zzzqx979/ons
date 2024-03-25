package test

import (
	"fmt"
	"github.com/spf13/viper"
	"ons/dataPanel/core"
	"ons/db"
	"ons/router"
	"testing"
)

func TestInitAIotServer(t *testing.T) {
	viper.SetConfigFile("../aiot.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal config file errors: %w", err))
	}
	db.DatabaseCon()
	db.RedisCon()
	router.InitAIotHttpRouter()
}

func TestQueryUpstream(t *testing.T) {
	viper.SetConfigFile("../gateway1.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal config file errors: %w", err))
	}
	t.Log(core.QueryOnsUpstream("string"))
	//t.Log(core.QueryOnsUpstream(core.GatewayOnsQueryByEpc, "test_epc1"))
	//t.Log(core.QueryOnsUpstream(core.GatewayOnsQueryByUrl, "/uv/321"))
}
