protoc-install:
	apt install -y protobuf-compiler

grpc-gen-server:
	protoc -I . \
       --go_out ./gen/go/ --go_opt paths=source_relative \
       --go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative \
       api/v1/hello_world.proto
	protoc -I . --grpc-gateway_out ./gen/go \
	 --grpc-gateway_opt logtostderr=true \
	 --grpc-gateway_opt paths=source_relative \
	 --grpc-gateway_opt generate_unbound_methods=true \
	 api/v1/hello_world.proto
