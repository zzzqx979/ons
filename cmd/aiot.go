package cmd

import (
	"github.com/spf13/cobra"
	"ons/dataPanel/core"
	"ons/dataPanel/mqtt"
	"ons/db"
	"ons/router"
)

var aIotCmd = &cobra.Command{
	Use:   "aiot",
	Short: "aiot tid server",
	Run:   RunAIot,
}

func RunAIot(cmd *cobra.Command, args []string) {
	InitConfig()
	db.DatabaseCon()
	db.SyncAIotTables()
	db.RedisCon()
	core.InitAIotOM()
	mqtt.InitMqttClient()
	router.InitAIotHttpRouter()
}
