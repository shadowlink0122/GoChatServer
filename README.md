# GoChatServer
Go言語で書いた、コマンドラインでチャットをするプログラム

GuestOS側でサーバを起動し、HostOS側でClientを起動して使います。


# Vagrantの準備

	brew install Vagrant

	vagrant up

# ビルド方法(Server)

	make server

# ビルド方法(Client)

	make client

# 起動方法(Server)

	vagrant ssh

	cd /vagrant/

	make run_s

# 起動方法(Client)
  
	make run_c
