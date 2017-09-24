# Strukture IDE Beta
[![Build Status](https://travis-ci.org/thestrukture/IDE.svg?branch=master)](https://travis-ci.org/thestrukture/IDE)
Go lang IDE server. Built with [Go Server](http://golangserver.com)

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
- NodeJS. (Build local ui) [Downloads](https://nodejs.org/en/download/)
- [Golang server](http://golangserver.com) if using `makefile`.
 


## Install

		$ go get github.com/thestrukture/IDE

#### How to run

		$ IDE

#### Launch with GUI
	
		$ make uiview
		$ IDE --headless & npm start

Take note of the pid ID to stop server process. Once your server is up feel free to use `npm start` directly.
	
	
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

## Community
Access the Strukture forums [here](http://forum.golangserver.com/forumdisplay.php?fid=3)

## Wiki : How to use the strukture

Visit https://github.com/thestrukture/IDE/wiki


## Debugging
Please use current golang command line tools as a fallback. The final result of each [Go server](http://golangserver.com) webapp is go code.

### Go Server apps
- [x] Find bugs by reading your web application's output.
If you're writing a webapp, Go Server attempts to use logging to find bugs. This process works by recovering from a runtime panic and stating the defunct line including the reason it crashed. This logger is only supported within the web service code, template pipeline code and templates. Due to this feature limit your code within the scope of one line.

Good :

	var name := Struct{Property:"val"}
	
Bad for debug logger :

	var name := Struct{ 
			    Property:"val",
			    FieldTwo: 2
		  }

#### How to recreate data for test
Use the `test` section within each package tree on the Strukture. Within this panel you may test pipelines, services as well as templates. 

## How to automate :
To build additional functionality use shell scripts. The users of your plugin can simply run the shell script via the HTTP terminal.

## Misc info
How to install GoS incase the built-in installer fails.

- Install GoS [CLI](http://golangserver.com). ( `$ go get github.com/cheikhshift/gos` )
- Install `GoS` dependencies : `$ gos deps`

# Contributions
Improvements to the codebase and pull requests are encouraged.
