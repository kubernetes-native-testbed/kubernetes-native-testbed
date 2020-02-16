COMMIT_HASH=`git rev-parse --short HEAD`

proto:
	protoc ./protobuf/comment.proto -I ~/protocolbuffers/protobuf/src -I ./protobuf --go_out=plugins=grpc:./protobuf --js_out=import_style=commonjs:../admin/src/js/protobuf --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../admin/src/js/protobuf

build:
	docker build -t k8stestbed/comment:$(COMMIT_HASH) .

push: build
	docker push k8stestbed/comment:$(COMMIT_HASH)


.PHONY: proto build push
