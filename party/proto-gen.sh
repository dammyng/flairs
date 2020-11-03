#protoc --proto_path=api/proto/v1 --proto_path=party --go_out=plugins=grpc:pkg/api/v1 flairs-service.proto
#protoc --proto_path=api/proto/v1 --proto_path=party --go_out=plugins=grpc:auth/pkg/api/v1 flairs-service.proto


protoc --proto_path=auth/api/proto/v1 --proto_path=party --go_out=plugins=grpc:auth/pkg/api/v1 flairs-service.proto
protoc --proto_path=auth/api/proto/v1 --proto_path=party --grpc-gateway_out=logtostderr=true:auth/pkg/api/v1 flairs-service.proto
protoc --proto_path=auth/api/proto/v1 --proto_path=party --swagger_out=logtostderr=true:auth/api/swagger/v1 flairs-service.proto


protoc --proto_path=wallet/api/proto/v1 --proto_path=party --go_out=plugins=grpc:wallet/pkg/api/v1 flairs-wallet.proto
protoc --proto_path=wallet/api/proto/v1 --proto_path=party --grpc-gateway_out=logtostderr=true:wallet/pkg/api/v1 flairs-wallet.proto
protoc --proto_path=wallet/api/proto/v1 --proto_path=party --swagger_out=logtostderr=true:wallet/api/swagger/v1 flairs-wallet.proto


protoc --proto_path=transaction/api/proto/v1 --proto_path=party --go_out=plugins=grpc:transaction/pkg/api/v1 flairs-transaction.proto
protoc --proto_path=transaction/api/proto/v1 --proto_path=party --grpc-gateway_out=logtostderr=true:transaction/pkg/api/v1 flairs-transaction.proto
protoc --proto_path=transaction/api/proto/v1 --proto_path=party --swagger_out=logtostderr=true:transaction/api/swagger/v1 flairs-transaction.proto


$SHELL
