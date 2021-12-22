gen-proto:
	protoc  --proto_path=proto/v1 --proto_path=third_party  \
					--go_out ./proto/v1 --go_opt paths=source_relative \
					--go-grpc_out ./proto/v1 --go-grpc_opt paths=source_relative \
					./proto/v1/todos.proto

gen-proto-gateway:
	protoc 	--proto_path=proto/v1 --proto_path=third_party --grpc-gateway_out ./proto/v1 \
					--grpc-gateway_opt logtostderr=true \
					--grpc-gateway_opt paths=source_relative \
					./proto/v1/todos.proto


