run:
	cd cmd/api && go run main.go

proto booking:
	cd pkg/pb && protoc --go_out=. --go-grpc_out=. booking.proto

proto2 admin:
	cd pkg/pb && protoc --go_out=. --go-grpc_out=. admin.proto