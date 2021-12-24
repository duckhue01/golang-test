gen-proto:
	protoc  --proto_path=proto/v2 --proto_path=third_party  \
					--go_out ./proto/v2 --go_opt paths=source_relative \
					--go-grpc_out ./proto/v2 --go-grpc_opt paths=source_relative \
					./proto/v2/todos.proto

gen-proto-gateway:
	protoc 	--proto_path=proto/v2 --proto_path=third_party --grpc-gateway_out ./proto/v2 \
					--grpc-gateway_opt logtostderr=true \
					--grpc-gateway_opt paths=source_relative \
					./proto/v2/todos.proto


