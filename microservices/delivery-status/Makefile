proto:
	protoc ./protobuf/delivery_status.proto -I ~/work/protocolbuffers/protobuf/src -I ./protobuf --go_out=plugins=grpc:./protobuf --js_out=import_style=commonjs:../admin/src/js/protobuf --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../admin/src/js/protobuf

.PHONY: proto
