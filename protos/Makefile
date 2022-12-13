GO_GRPC=../src/go
TS_SERVER=../ts-server
JS_CLIENT=../js-client-vue
TS_REACT=../ts-client-react
TS_CLIENT=../ts-client
ENVOY_DIR=../envoy
PROTO_DIR=.

# git clone https://github.com/googleapis/googleapis
GOOGLEAPIS_DIR=../../googleapis/googleapis

all: clean go-grpc

clean:
	rm -f $(GO_GRPC)/*.pb.go 
	rm -f $(JS_CLIENT)/src/models/todo_grpc_web_pb.js $(JS_CLIENT)/src/models/todo_pb.js
	rm -f $(TS_REACT)/src/models/todo_grpc_web_pb.js $(TS_REACT)/src/models/todo_pb.js
	rm -f $(TS_SERVER)/src/grpcjs/todo_grpc_web_pb.js $(TS_SERVER)/src/grpcjs/todo_pb.js

gen: envoy go-grpc-todo go-grpc-auth ts-server js-client-vue ts-client-react ts-client

go-grpc:
	#
	# Go Server
	#
	
	
	# gRPC versions
#	protoc \
#		-I . \
#		--proto_path=$(PROTO_DIR) \
#		--go_out=$(GO_GRPC)/genpb --go_opt=paths=source_relative \
#		--twirp_out=$(GO_GRPC)/twpb/publicapi \
#		--go-grpc_out=$(GO_GRPC)/genpb --go-grpc_opt=paths=source_relative \
#		./publicapi/*
#	protoc \
#		-I . \
#		--proto_path=$(PROTO_DIR) \
#		--go_out=$(GO_GRPC)/genpb --go_opt=paths=source_relative \
#		--twirp_out=$(GO_GRPC)/twpb/core \
#		--go-grpc_out=$(GO_GRPC)/genpb --go-grpc_opt=paths=source_relative \
#		./core/*

	# Twirp versions
	protoc \
		-I . \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_GRPC)/twpb --go_opt=paths=source_relative \
		--twirp_out=$(GO_GRPC)/twpb/publicapi \
		./publicapi/*
	protoc \
		-I . \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_GRPC)/twpb --go_opt=paths=source_relative \
		--twirp_out=$(GO_GRPC)/twpb/core \
		./core/*


ts-server:
	protoc \
		-I . -I ./protos \
		--proto_path=$(PROTO_DIR) \
		--plugin=protoc-gen-ts=$(TS_SERVER)/node_modules/.bin/protoc-gen-ts \
		--ts_out=grpc_js:$(TS_SERVER)/src/grpcjs \
		./todo.proto google/api/*.proto
#
# JS Client
#
js-client-vue:
	protoc \
		-I . -I ./protos \
		--proto_path=$(PROTO_DIR) \
		--js_out=import_style=commonjs,binary:$(JS_CLIENT)/models/ \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:$(JS_CLIENT)/models/ \
		./todo.proto

#
# TS React Client
#
ts-client-react:
	#protoc \
	#	-I . \
	#	--proto_path=$(PROTO_DIR) \
	#	--plugin=$(TS_REACT)/node_modules/.bin/protoc-gen-ts_proto \
	#	--ts_proto_opt=esModuleInterop=true \
	#	--ts_proto_opt=outputClientImpl=grpc-web \
	#	--ts_proto_out=import_style=commonjs,binary:$(TS_REACT)/src/models/ \
	#	--plugin="protoc-gen-ts=${TS_REACT}/node_modules/.bin/protoc-gen-ts" \
	#	publicapi/*
	protoc \
		-I . \
		--proto_path=$(PROTO_DIR) \
		--grpc-web_out=import_style=commonjs+dts,mode=grpcweb:$(TS_REACT)/src/genpb \
	    --js_out=import_style=commonjs,binary:$(TS_REACT)/src/genpb \
		publicapi/* google/rpc/* google/api/*

#
# TS Command Line Client
#
#protoc 

ts-client:
	# --plugin=protoc-gen-ts=$(TS_CLIENT)/node_modules/protoc-gen-ts/bin/protoc-gen-ts \

	protoc \
		-I . -I ./protos \
		--proto_path=$(PROTO_DIR) \
		--plugin=protoc-gen-ts=/Users/koblas/repos/github.com/koblas/protoc-gen-ts/bin/protoc-gen-ts \
		--ts_out=grpc_js:$(TS_CLIENT)/src/grpcjs \
		todo.proto google/api/*.proto

#
# Envoy
#
envoy:
	protoc \
		-I . \
		--proto_path=$(PROTO_DIR) \
		--include_imports --include_source_info \
		--descriptor_set_out=$(ENVOY_DIR)/common.pb \
		publicapi/* google/rpc/error_details.proto
