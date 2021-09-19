GO_SERVER=./go-server
TS_SERVER=./ts-server
JS_CLIENT=./js-client-vue
TS_CLIENT=./ts-client-react
PROTO_DIR=./protos

all: clean gen

clean:
	rm -f todo/todo.pb.go todo/todo_grpc.pb.go
	rm -f todo-client/models/todo_grpc_web_pb.js todo-client/models/todo_pb.js
	rm -f todo-ts-client/models/todo_grpc_web_pb.js todo-ts-client/models/todo_pb.js

gen:
	# Go Server
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(GO_SERVER)/todo --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_SERVER)/todo --go-grpc_opt=paths=source_relative \
		protos/todo.proto
	# TS Server
	#     https://github.com/agreatfool/grpc_tools_node_protoc_ts/blob/master/examples/bash/build.sh
	
	$(TS_SERVER)/node_modules/.bin/grpc_tools_node_protoc \
		--js_out=import_style=commonjs,binary:$(TS_SERVER)/src/grpcjs \
		--grpc_out=grpc_js:$(TS_SERVER)/src/grpcjs \
		-I ./protos \
		protos/todo.proto
	$(TS_SERVER)/node_modules/.bin/grpc_tools_node_protoc \
		--plugin=protoc-gen-ts=$(TS_SERVER)/node_modules/.bin/protoc-gen-ts \
		--ts_out=grpc_js:$(TS_SERVER)/src/grpcjs \
		-I ./protos \
		protos/todo.proto
	protoc --proto_path=$(PROTO_DIR) \
		--ts_proto_opt=outputServices=grpc-js \
		--ts_proto_opt=esModuleInterop=true \
		--plugin=$(TS_SERVER)/node_modules/.bin/protoc-gen-ts_proto \
		--ts_proto_out=import_style=commonjs,binary:$(TS_SERVER)/src/grpcjs/ \
		protos/todo.proto


	# JS Client
	protoc --proto_path=$(PROTO_DIR) \
		--js_out=import_style=commonjs,binary:$(JS_CLIENT)/models/ \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:$(JS_CLIENT)/models/ \
		protos/todo.proto
	# TS Client
	protoc --proto_path=$(PROTO_DIR) \
		--ts_proto_opt=outputClientImpl=grpc-web \
		--plugin=$(TS_CLIENT)/node_modules/.bin/protoc-gen-ts_proto \
		--ts_proto_out=import_style=commonjs,binary:$(TS_CLIENT)/src/models/ \
		protos/todo.proto
