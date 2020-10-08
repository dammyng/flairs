#protoc --proto_path=api/proto/v1 --proto_path=party --go_out=plugins=grpc:pkg/api/v1 flairs-service.proto
#protoc --proto_path=api/proto/v1 --proto_path=party --go_out=plugins=grpc:auth/pkg/api/v1 flairs-service.proto


protoc --proto_path=api/proto/v1 --proto_path=party --go_out=plugins=grpc:auth/pkg/api/v1 flairs-service.proto
protoc --proto_path=api/proto/v1 --proto_path=party --grpc-gateway_out=logtostderr=true:auth/pkg/api/v1 flairs-service.proto
protoc --proto_path=api/proto/v1 --proto_path=party --swagger_out=logtostderr=true:api/swagger/v1 flairs-service.proto

$SHELL
