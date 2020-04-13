#!/usr/bin/make -f
# ----------------------------------------------------------------------
#
generate-mail-go:
	mkdir -p ./generated/go/mail
	protoc  \
	  --proto_path=${GOPATH}/src \
	  --proto_path=./proto/mail \
	  --go_out=generated/go/mail \
	 ./proto/mail/*.proto

generate-detection-go:
	mkdir -p ./generated/go/detection
	protoc  \
	  --proto_path=${GOPATH}/src \
	  --proto_path=./proto/mail \
	  --proto_path=./proto/detection \
	  --go_out=:generated/go/detection \
	 ./proto/detection/*.proto

generate-java:
	./gradlew clean build

generate-go: generate-mail-go generate-detection-go
