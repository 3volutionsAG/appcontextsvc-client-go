# protoc --go_out=clients/go --go_opt=paths=source_relative \
# 		--experimental_allow_proto3_optional=true \
# 		-I=src/Protos \
#     src/Protos/**/*.proto src/Protos/*.proto

protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	--experimental_allow_proto3_optional=true \
	-I=protos protos/*.proto protos/**/*.proto 

go mod tidy