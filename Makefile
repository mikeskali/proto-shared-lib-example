#!/usr/bin/make -f
# ----------------------------------------------------------------------
#
generate-java-protoc:
	mkdir -p ./generated/java/mail ./generated/java/detection
	protoc  \
    	  --proto_path=${GOPATH}/src \
    	  --proto_path=./proto/mail \
    	  --proto_path=./proto/detection \
    	  --java_out=:generated/java/ \
    	 ./proto/mail/*.proto ./proto/detection/*.proto

generate-go-protoc:
	mkdir -p ./generated/go/mail ./generated/go/detection
	for x in "mail" "detection"; do \
		protoc  \
    	  --proto_path=./proto/mail \
    	  --proto_path=./proto/detection \
    	  --go_out=:generated/go/$$x \
    	 ./proto/$$x/*.proto; \
    	done

generate-protoc: generate-java-protoc generate-go-protoc


generate-java-gradle:
	./gradlew clean build

generate-go: generate-mail-go generate-detection-go
