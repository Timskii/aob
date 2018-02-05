package main

import (
		"fmt"
		

		"github.com/spf13/cobra"

		"ps.kz/aob/abac/server"
)

var mainCmd = &cobra.Command{
	Use: "main",
	Run: func(cmd *cobra.Command, args []string){
			fmt.Println("test")
		},
}

func main() {

	mainCmd.AddCommand(server.Cmd())

	mainCmd.Execute()
}