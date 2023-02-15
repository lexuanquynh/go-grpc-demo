proto:
  protoc --go_out=src/go/sma --go_opt=paths=source_relative \
	 --go-grpc_out=src/go/sma --go-grpc_opt=paths=source_relative \
	 --grpc-gateway_out=logtostderr=true,paths=source_relative, \
	   grpc_api_configuration=$(annotation_file):src/go/sma \
	 --swagger_out=grpc_api_configuration=$(annotation_file):src/swagger \
	 --openapiv2_out=logtostderr=true:src/openapi $(proto_file)