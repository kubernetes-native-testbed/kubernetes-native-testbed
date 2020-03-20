COMMIT_HASH=`git rev-parse --short HEAD`

proto:
	protoc ./protobuf/product.proto -I ./protobuf --go_out=plugins=grpc:./protobuf --js_out=import_style=commonjs:../admin/src/js/protobuf --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../admin/src/js/protobuf

build:
	docker build -t k8stestbed/product:$(COMMIT_HASH) .

push: build
	docker push k8stestbed/product:$(COMMIT_HASH)


.PHONY: proto build push

