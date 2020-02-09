COMMIT_HASH=`git rev-parse --short HEAD`

build:
	docker build -t k8stestbed/admin:$(COMMIT_HASH) .

push: build
	docker push k8stestbed/admin:$(COMMIT_HASH)

.PHONY: build push
