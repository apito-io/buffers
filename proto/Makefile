default: proto

.PHONY: proto

proto:
		protoc --go_out=$(GOPATH)/src --go_opt=paths=import --go-grpc_out=$(GOPATH)/src --go-grpc_opt=paths=import codegen.proto

