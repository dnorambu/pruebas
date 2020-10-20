.PHONY: compile

compile:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative courier/courier.proto
cliente:
	go run cliente_courier/cliente.go
camion:
	go run camion_courier/camion.go
logistica:
	go run server_courier/server.go