# Strukture IDE Beta
[![Build Status](https://travis-ci.org/thestrukture/IDE.svg?branch=master)](https://travis-ci.org/thestrukture/IDE)
Go lang IDE server.

## About project
The strukture is an open source IDE. It is designed to build web applications, with the organizational help of [Go-Server](http://golangserver.com). The IDE runs as a server and is accessed via web browser. Being a web server, the IDE is accessible from any device on your network. Compared to Microsoft's VS Code and Eclipse CHE this IDE is very minimal. It features :
- Web application resource management.
- Autocomplete between different files.
- Syntax correction.
- Server process management.
- Basic terminal via stateless HTTP.
- Line tags in relation to failed build logs.
- Project build scripts.
- View web application output.

## Requirements
- Unix based OS/CLI.
- Go 1.9 and up. [Find it here](https://golang.org/dl/).




## Install

		$ go get github.com/thestrukture/IDE


### Install Via APT

#### Ubuntu 14.04
	
	wget -qO- https://dl.packager.io/srv/thestrukture/IDE/key | sudo apt-key add -
	sudo wget -O /etc/apt/sources.list.d/ide.list \
  	https://default:660f1018d50549e01c9d6f3a6d8f66c8@dl.packager.io/srv/thestrukture/IDE/master/installer/ubuntu/14.04.repo
	sudo apt-get update
	sudo apt-get install ide

#### Ubuntu 16.04

	wget -qO- https://dl.packager.io/srv/thestrukture/IDE/key | sudo apt-key add -
	sudo wget -O /etc/apt/sources.list.d/ide.list \
  	https://default:660f1018d50549e01c9d6f3a6d8f66c8@dl.packager.io/srv/thestrukture/IDE/master/installer/ubuntu/16.04.repo
	sudo apt-get update
	sudo apt-get install ide
	
#### How to run

	$ ide
	
[Link to packager.io](https://packager.io/gh/thestrukture/IDE)

## Run

		$ IDE

## Access

Visit [localhost:8884/home](http://localhost:8884/home). Access the IDE from any device on your network as well...

## Bug reports & questions :
Please create a new issue on Github to report a bug.

## Community
Access the Strukture forums [here](http://forum.golangserver.com/forumdisplay.php?fid=3)

## Wiki : How to use the strukture

Visit https://github.com/thestrukture/IDE/wiki

## How to automate :
To build additional functionality use shell scripts. The users of your plugin can simply run the shell script via the HTTP terminal.

## Misc info
How to install GoS incase the built-in installer fails.

- Install GoS [CLI](http://golangserver.com). ( `$ go get github.com/cheikhshift/gos` )
- Install `GoS` dependencies : `$ gos deps`

## Screenshots

![](tests/1.png)
![](tests/8.png)
![](tests/30.png)
![](tests/29.png)
![](tests/20.png)
![](tests/2.png)
![](tests/24.png)
![](tests/3.png)
![](tests/4.png)
![](tests/5.png)
![](tests/25.png)
![](tests/6.png)
![](tests/7.png)

