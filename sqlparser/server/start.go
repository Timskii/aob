package server

import (
		
	"fmt"
	"net"

	"github.com/spf13/cobra"
)

func startCmd() *cobra.Command {
	return serverStartCmd
}

var serverStartCmd = &cobra.Command{
	Use: "start",
	/*Short: "",
	Long: "",*/
	RunE: func(cmd *cobra.Command, args []string) error {
		return startServer(args)
	},
}


func startServer(args []string)  error {
	lis, err := net.Listen("tcp","localhost:4001")
	if err!=nil {
		fmt.Println("err =", err)
	}
	defer lis.Close()
	fmt.Println("server srarted")
	return nil

}