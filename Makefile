#!/usr/bin/make -f
# ----------------------------------------------------------------------
#
generate-java-protoc:
	mkdir -p ./generated/java/mail ./generated/java/detection
	protoc  \
    	  --proto_path=./proto/ \
    	  --proto_path=./proto/validator/ \
    	  --java_out=:generated/java/ \
    	 ./proto/mail/*.proto ./proto/detection/*.proto

generate-go-protoc:
	mkdir -p ./generated/go/mail ./generated/go/detection
	go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
	for x in "mail" "detection"; do \
		protoc  \
    	  --proto_path=${GOPATH}/src \
		  --proto_path=${GOPATH}/src/github.com/google/protobuf/src \
		  --proto_path=${GOPATH}/src/github.com/mwitkow/go-proto-validators \
    	  --proto_path=./proto \
    	  --go_out=plugins=grpc,paths=source_relative:generated/go/ \
		  --govalidators_out=paths=source_relative:generated/go/ \
    	 ./proto/$$x/*.proto; \
    	done



generate-java-gradle:
	./gradlew clean build

generate-protoc: generate-java-protoc generate-go-protoc