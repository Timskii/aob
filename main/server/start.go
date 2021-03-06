package server

import (
		
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	pb "ps.kz/aob/main/protos"
	"ps.kz/aob/main/core"
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
	lis, err := net.Listen("tcp","localhost:4000")
	if err!=nil {
		log.Println("err =", err)
	}
	defer lis.Close()

	ser := grpc.NewServer()
	pb.RegisterAOBServiceServer(ser, core.NewAdminServer())
	ser.Serve(lis)
	log.Println("server srarted")
	return nil

}