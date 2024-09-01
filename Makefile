
test:
	go test -v ./... -coverpkg=./...
gen:
	protoc -I ./pb  --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=true --govalidators_out=. \
		product-sub.proto
	mockery --all --recursive --dir ./pkg
	mockery --dir ./pb  --all --recursive --output ./pb/gen/mocks

deps:
	go get google.golang.org/protobuf/cmd/protoc-gen-go \
			google.golang.org/grpc/cmd/protoc-gen-go-grpc \
			github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
			github.com/vektra/mockery/v2/.../
