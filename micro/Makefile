default: proto

.PHONY: proto

proto:
		protoc --go_out=$(GOPATH)/src --go_opt=paths=import --micro_out=$(GOPATH)/src schema.proto

