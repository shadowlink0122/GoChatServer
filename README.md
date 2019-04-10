# GoChatServer
ChatServer by Golang

プログラム中の "IPAdress:Port" の部分を、あなたの環境にあったものに変更して動作させてください。

# 起動方法
Server側
  vagrant ssh
  cd /vagrant/go_app/tcp
  go run server.go

Client側
  cd go_app/tcp
  go run client.go
