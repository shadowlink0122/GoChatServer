# GoChatServer
ChatServer by Golang

プログラム中の "IPAdress:Port" の部分を、あなたの環境にあったものに変更して動作させてください。

# ビルド方法(Server)

ゲストOS側で実行

	vagrant ssh

	cd /vagrant

	make server

# ビルド方法(Client)

ホストOS側で実行

	make client

# 起動方法(Server)

先にサーバの起動をしてください。

	vagrant ssh

	cd /vagrant

	make run_s

# 起動方法(Client)
  
ホストOS側で実行

	make run_c
