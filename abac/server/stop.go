package server

import (
		"fmt"

		"github.com/spf13/cobra"
)

func stopCmd() *cobra.Command{
	return serverStopCmd
}

var serverStopCmd = &cobra.Command{
	Use: "stop",
	RunE: func(cmd *cobra.Command, args []string) error{
		return stop()
	},
}

func stop() error {
	fmt.Println("server is stoped")
	return nil
}