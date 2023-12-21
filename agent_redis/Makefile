print:
	@echo "Hello"

protobuf:
	@protoc -I ../protobuf/agent --go_out=./pkg snap.proto
	@protoc -I ../protobuf/agent/redis --go_out=./pkg client_list.proto
	@protoc -I ../protobuf/agent/redis --go_out=./pkg db_size.proto
	@protoc -I ../protobuf/agent/redis --go_out=./pkg info_cpu.proto
	@protoc -I ../protobuf/agent/redis --go_out=./pkg info_memory.proto
	@protoc -I ../protobuf/agent/redis --go_out=./pkg info_stat.proto
	@protoc -I ../protobuf/agent/redis --go_out=./pkg redis_snap.proto

build:
	@go -o ./dist/dist/agent_redis ./cmd/execute/*

test:
	@go test -v ./pkg/db/

all: protobuf build