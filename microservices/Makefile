COMMIT_HASH=`git rev-parse --short HEAD`
DOCKER_BUILD=docker build .
DOCKER_PUSH=docker push
IMAGE_PATH=k8stestbed
BUILD_ARG=--build-arg MICROSERVICE

admin:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

adminp: admin
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

cart:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

cartp: cart
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

comment:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

commentp: comment
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

order:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

orderp: order
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

delivery-status:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

delivery-statusp: delivery-status
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

payment-info:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

payment-infop: payment-info
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

point:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

pointp: point
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

product:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

productp: product
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

rate:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

ratep: rate
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

user:
	$(DOCKER_BUILD) -f $@/Dockerfile -t $(IMAGE_PATH)/$@:$(COMMIT_HASH)

userp: user
	$(DOCKER_PUSH) $(IMAGE_PATH)/$<:$(COMMIT_HASH)

.PHONY: admin adminp cart cartp order orderp comment commentp delivery-status delivery-statusp payment-info payment-infop point pointp product productp rate ratep user userp

