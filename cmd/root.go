package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var configFile string

var RootCmd = &cobra.Command{
	Use:   "go-starter",
	Short: "An opinionated Go starter kit built on top of Chi",

	Run: func(cmd *cobra.Command, args []string) {
		serverType := strings.ToLower(os.Getenv("SERVER_TYPE"))
		switch serverType {
		case "http":
			cmd.Run(HTTPCmd, args)
		default:
			panic(fmt.Sprintf("server type '%s' is not supported", serverType))
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.auth.yaml)")
}
