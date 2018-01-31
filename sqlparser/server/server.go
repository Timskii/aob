package server

import (

	"github.com/spf13/cobra"
)


func Cmd() *cobra.Command {
	serverCmd.AddCommand(startCmd())
	serverCmd.AddCommand(stopCmd())
	return serverCmd
}

var serverCmd = &cobra.Command{
	Use: "server",
}

