server:
	go run ./go_app/tcp/server.go ./go_app/tcp/encrypt.go ./go_app/tcp/decrypt.go 

client:
	go run ./go_app/tcp/client.go ./go_app/tcp/encrypt.go ./go_app/tcp/decrypt.go

