# Strukture IDE Beta
[![Build Status](https://travis-ci.org/thestrukture/IDE.svg?branch=master)](https://travis-ci.org/thestrukture/IDE)
Go lang IDE server. Built with [Go Server](http://gophersauce.com)

## About project
The strukture is an open source IDE. It is designed to build web applications, with the organizational help of [Go-Server](http://gophersauce.com). The IDE runs as a server and is accessed via web browser. Being a web server, the IDE is accessible from any device on your network. Compared to Microsoft's VS Code and Eclipse CHE this IDE is very minimal. It features :
- Web application resource management.
- Autocomplete between different files.
- Syntax correction.
- Server process management.
- Basic terminal via stateless HTTP.
- Line tags in relation to failed build logs.
- Project build scripts.
- View web application output.

## Requirements
- Go 1.9 and up. [Find it here](https://golang.org/dl/).
- [Golang server](http://gophersauce.com) if using `makefile`.

### If plan on using Electron View. A node JS project :
- NodeJS. (Build local ui) [Downloads](https://nodejs.org/en/download/)


## Install

		$ go get github.com/thestrukture/IDE

#### How to run

		$ IDE

#### Launch with GUI (Electron View)
Change to the ui directory within the root of this package. (Moved ui files to build package as go)

		$ cd ui/
		$ npm install
		$ IDE --headless & npm start

Take note of the pid ID to stop server process. Once your server is up feel free to use `npm start` directly.

### IE Fix
If build commands keep returning the same message, push the `F12` key down to open developer tools and try building again.
	
### Install Via APT

#### Ubuntu 14.04
	
	wget -qO- https://dl.packager.io/srv/thestrukture/IDE/key | sudo apt-key add -
	sudo wget -O /etc/apt/sources.list.d/ide.list \
  	https://dl.packager.io/srv/thestrukture/IDE/master/installer/ubuntu/14.04.repo
	sudo apt-get update
	sudo apt-get install ide

#### Ubuntu 16.04

	wget -qO- https://dl.packager.io/srv/thestrukture/IDE/key | sudo apt-key add -
	sudo wget -O /etc/apt/sources.list.d/ide.list \
  	https://dl.packager.io/srv/thestrukture/IDE/master/installer/ubuntu/16.04.repo
	sudo apt-get update
	sudo apt-get install ide
	
#### How to run

	$ ide
	
[Link to packager.io](https://packager.io/gh/thestrukture/IDE)


## Access

Visit [localhost:8884/home](http://localhost:8884/home). Access the IDE from any device on your network as well...

## Bug reports & questions :
Please create a new issue on Github to report a bug.

## Wiki : How to use the strukture

Visit https://github.com/thestrukture/IDE/wiki

## Misc info
How to install GoS incase the built-in installer fails.

- Install GoS [CLI](http://gophersauce.com). ( `$ go get github.com/cheikhshift/gos` )
- Install `GoS` dependencies : `$ gos deps`

# Contributions
Improvements to the codebase and pull requests are encouraged.

## Screenshots

![screenshot](tests/1newsc.png)
![screenshot](tests/2newsc.png)
![screenshot](tests/3newsc.png)
![screenshot](tests/4newsc.png)
![screenshot](tests/5newsc.png)
