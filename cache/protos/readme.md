#readme


go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc

example
protoc --go_out=plugins=grpc:. *.proto