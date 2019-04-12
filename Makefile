server:
	go build ./go_app/tcp/server.go ./go_app/tcp/encrypt.go ./go_app/tcp/decrypt.go 
	mv server ./go_app

client:
	go build ./go_app/tcp/client.go ./go_app/tcp/encrypt.go ./go_app/tcp/decrypt.go
	mv client ./go_app

run_s:
	./go_app/server

run_c:
	./go_app/client

