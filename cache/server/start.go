package server

import (
		
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	pb "ps.kz/aob/cache/protos"
	"ps.kz/aob/cache/core"
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
	lis, err := net.Listen("tcp","localhost:4002")
	if err!=nil {
		fmt.Println("err =", err)
	}
	defer lis.Close()

	ser := grpc.NewServer()
	pb.RegisterAOBServiceServer(ser, core.NewAdminServer())
	ser.Serve(lis)
	fmt.Println("server srarted")
	return nil

}