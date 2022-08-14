gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:.   --grpc-gateway_out=:. --openapiv2_out=:swagger

grpcServer:
	go run ./server/main.go -port 8080

clientServer:
	go run ./client/main.go -address 0.0.0.0:8080

unitTest:
	cd service && go test -covermode=count -coverprofile=coverprofile.cov -run="^Test"