proto-generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./Protos/database/mongoclient/mongoclient.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./Protos/cart/cart.proto

tidy:
	go mod tidy

mongoclient:
	go run Servers/mongoclient/main.go
	
alldockerup:
	sudo docker-compose up -d

buildallservice:
	docker build --target apigateway -t apigateway:latest .
	docker build --target cartserver -t cartserver:latest .
	docker build --target mongoclientserver -t mongoclientserver:latest .

buildpods:
	kubectl create -f k8s/elastic.yaml