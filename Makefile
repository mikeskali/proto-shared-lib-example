#!/usr/bin/make -f
# ----------------------------------------------------------------------
#
generate-java-protoc:
	mkdir -p ./generated/java/mail ./generated/java/detection
	protoc  \
    	  --proto_path=./proto/ \
    	  --java_out=:generated/java/ \
    	 ./proto/mail/*.proto ./proto/detection/*.proto

generate-go-protoc:
	mkdir -p ./generated/go/mail ./generated/go/detection
	for x in "mail" "detection"; do \
		protoc  \
    	  --proto_path=./proto \
    	  --go_out=paths=source_relative:generated/go/ \
    	 ./proto/$$x/*.proto; \
    	done


generate-java-gradle:
	./gradlew clean build

generate-protoc: generate-java-protoc generate-go-protoc