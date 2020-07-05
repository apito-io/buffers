COMMON_DIR?=common

default: sob-build-kor

.PHONY: sob-build-kor $(ALL) $(PROFILE_DIR) $(COMMON_DIR) $(SEARCH_DIR)

sob-build-kor:
		make common

$(ALL):
		make profile && make chat && make notification && make vendors && make place && make search && make graph

$(COMMON_DIR):
		protoc --gogofaster_out=plugins=grpc:$(GOPATH)/src $(COMMON_DIR)*.proto gogo.proto chunk.proto --dart_out=/Users/diablo/Projects/flutter/udhvabon/xbackend_web_mvp/lib/buffers