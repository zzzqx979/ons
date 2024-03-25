package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFileFlag string

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "uq valley tid sever.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("use '-h' echo program help.")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&cfgFileFlag, "config", "c", "app.yaml",
		"init tid sever config by config file.support type:yaml,json,env")

	rootCmd.AddCommand(gatewayCmd, aIotCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorf("execute root cmd errors:%+v.\n", err)
		os.Exit(-1)
	}
}

func InitConfig() {
	// 设置配置文件
	viper.SetConfigFile(cfgFileFlag)
	// 寻找配置文件并读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal config file errors: %w", err))
	}
}
