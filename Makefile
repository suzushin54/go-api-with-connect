.PHONY: setup
setup:
	make install
	make install-check
	go mod tidy
	#buf mod init # already created and modified

.PHONY: install
install:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

install-check:
	which buf grpcurl protoc-gen-go protoc-gen-connect-go
	# 🚀 If you see four paths, you're good to go!

buf:
	buf lint
	buf generate


#.PHONY: gen-my_service-docker
#gen-my_service-docker:
#	docker build -t protoc-gen:v23.1 ./docker/protoc-gen
#	docker run -it -v ${PWD}:/home --rm protoc-gen:v23.1 make protoc
#	goimports -w ./pb
#
#
#.PHONY: protoc
#protoc:
#	mkdir -p ./pb
#	protoc --go_out=./pb --go_opt=paths=source_relative \
#	--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
#	./my_service/*.my_service
