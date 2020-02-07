COMMIT_HASH=`git rev-parse --short HEAD`

proto:
	protoc ./protobuf/product.proto -I ./protobuf --go_out=plugins=grpc:./protobuf --js_out=import_style=commonjs:./client --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client

client:
	docker build -t k8stestbed/client:$(COMMIT_HASH) ./client
	docker push k8stestbed/client:$(COMMIT_HASH)

server:
	docker build -t k8stestbed/product:$(COMMIT_HASH) .
	docker push k8stestbed/product:$(COMMIT_HASH)

.PHONY: proto client server
