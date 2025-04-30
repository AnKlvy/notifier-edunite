gen:
	@protoc --proto_path=protobuf "protobuf/notifier.proto" --go_out=protobuf/gen_notifier --go_opt=paths=source_relative --go-grpc_out=protobuf/gen_notifier --go-grpc_opt=paths=source_relative
