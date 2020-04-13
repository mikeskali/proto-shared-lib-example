#!/usr/bin/make -f
# ----------------------------------------------------------------------
#
generate-mail-go:
	go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
	mkdir -p ./generated/go/mail
	protoc  \
	  --proto_path=${GOPATH}/src \
	  --proto_path=${GOPATH}/src/github.com/mwitkow/go-proto-validators \
	  --proto_path=./proto/mail \
	  --go_out=paths=source_relative,Mvalidator.proto=github.com/mwitkow/go-proto-validators:generated/go/mail \
	  --govalidators_out=paths=source_relative,Mvalidator.proto=github.com/mwitkow/go-proto-validators:generated/go/mail \
	 ./proto/mail/*.proto

generate-detection-go:
	go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
	mkdir -p ./generated/go/detection
	protoc  \
	  --proto_path=${GOPATH}/src \
	  --proto_path=${GOPATH}/src/github.com/mwitkow/go-proto-validators \
	  --proto_path=./proto/mail \
	  --proto_path=./proto/detection \
	  --go_out=plugins=grpc,paths=source_relative,Mvalidator.proto=github.com/mwitkow/go-proto-validators:generated/go/detection \
	  --govalidators_out=paths=source_relative,Mvalidator.proto=github.com/mwitkow/go-proto-validators:generated/go/detection \
	 ./proto/detection/*.proto

generate-go: generate-mail-go generate-detection-go
