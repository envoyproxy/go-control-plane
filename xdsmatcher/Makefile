include ../Makefile.common

.PHONY: proto
proto: test/proto/*
	protoc --proto_path=test/proto --go_out=test/proto --go_opt=paths=source_relative test.proto
