# GoChatServer
Go言語で書いた、コマンドラインでチャットをするプログラム

GuestOS側でサーバを起動し、HostOS側でClientを起動して使います。


# Vagrantの準備

	brew install Vagrant

	vagrant up

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
