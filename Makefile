SRC = ./go_app/tcp

server:
	go build ${SRC}/server.go ${SRC}/encrypt.go ${SRC}/decrypt.go ${SRC}/definition.go
	mv server ./go_app

client:
	go build ${SRC}/client.go ${SRC}/encrypt.go ${SRC}/decrypt.go ${SRC}/definition.go
	mv client ./go_app

run_s:
	./go_app/server

run_c:
	./go_app/client

dist:
	rm -rf ./go_app/server
	rm -rf ./go_app/client

git: dist
	git checkout miyajima
	git add .
	git commit -m "Chat Server and Client"
	git push --set-upstream origin miyajima


