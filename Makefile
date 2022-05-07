proto-generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./Protos/users/user.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./Protos/database/postgresclient/postgresclient.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./Protos/database/mongoclient/mongoclient.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./Protos/product/product.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./Protos/cart/cart.proto

tidy:
	go mod tidy

mongoclient:
	go run Servers/mongoclient/main.go

postgresclient:
	go run Servers/Postgresclient/main.go

alldockerup:
	sudo docker-compose up -d
