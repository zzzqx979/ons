package cmd

import (
	"github.com/spf13/cobra"
	"ons/dataPanel/core"
	"ons/db"
	"ons/router"
)

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "gateway tid server",
	Run:   RunGateway,
}

func RunGateway(cmd *cobra.Command, args []string) {
	InitConfig()
	db.DatabaseCon()
	db.SyncGatewayTables()
	core.InitGatewayOM()
	router.InitGatewayHttpServer()
}
