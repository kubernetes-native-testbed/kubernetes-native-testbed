#!/bin/bash

OS_TYPE=$(uname | tr '[A-Z]' '[a-z]')
GPRC_WEB_VERSION=1.0.7
SKAFFOLD_VERSION=1.5.0

# install developmen tools
brew install protobuf

go get -u github.com/golang/protobuf/protoc-gen-go

curl -sL https://github.com/grpc/grpc-web/releases/download/${GPRC_WEB_VERSION}/protoc-gen-grpc-web-${GPRC_WEB_VERSION}-${OS_TYPE}-x86_64 -o /usr/local/bin/protoc-gen-grpc-web
chmod 755 /usr/local/bin/protoc-gen-grpc-web


# install skaffold
if ! which skaffold; then
  curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v${SKAFFOLD_VERSION}/skaffold-${OS_TYPE}-amd64
  chmod +x skaffold
  sudo mv skaffold /usr/local/bin
fi


# install telepresence
if ! which telepresence; then
  brew cask install osxfuse
  brew install datawire/blackbird/telepresence
fi

if [ -z "${REMOTE_CONTEXT}" ]; then
  echo "Please input REMOTE_CONTEXT env var"
  read REMOTE_CONTEXT
  export REMOTE_CONTEXT
fi
if [ -z "${LOCAL_CONTEXT}" ]; then
  echo "Please input LOCAL_CONTEXT env var"
  read LOCAL_CONTEXT
  export LOCAL_CONTEXT
fi

