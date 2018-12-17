//Package api holds the types for the phraser gRPC service
//go:generate protoc -I . -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. phraser.proto
//go:generate protoc -I . -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. phraser.proto
package api
