# Strukture IDE Beta II
[![Build Status](https://travis-ci.org/thestrukture/IDE.svg?branch=master)](https://travis-ci.org/thestrukture/IDE)
[![GoDoc](https://godoc.org/github.com/thestrukture/IDE/api?status.svg)](https://godoc.org/github.com/thestrukture/IDE/api)
[![Test Coverage](https://api.codeclimate.com/v1/badges/d46b0bfb51e827632710/test_coverage)](https://codeclimate.com/github/thestrukture/IDE/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/d46b0bfb51e827632710/maintainability)](https://codeclimate.com/github/thestrukture/IDE/maintainability)

Go lang IDE server. Built with [GopherSauce](http://gophersauce.com)

## About project
The strukture is an open source IDE. It is designed to build web applications, with the organizational help of [GopherSauce](http://gophersauce.com). The IDE runs as a server and is accessed via web browser. Being a web server, the IDE is accessible from any device on your network.Compared to Microsoft VS Code and Eclipse CHE, this IDE is very minimalistic. It features :
- Web application resource management.
- Autocomplete between different files.
- Syntax correction.
- Server process management.
- Basic terminal via stateless HTTP.
- Line tags in relation to failed build logs.
- Project build scripts.
- View web application output.
- Build docker images (Must have docker running on host) . 

## Requirements
- Go v1.7+.
- [GopherSauce](http://gophersauce.com) if using `makefile`.

## Install

		$ go get github.com/thestrukture/IDE

#### How to run

		$ IDE

#### Launch with GUI (Electron View)

###### Requires NodeJS
- NodeJS [Downloads](https://nodejs.org/en/download/)

Change to the ui directory within the root of this package. (Moved ui files to build package as go)

		$ cd ui/
		$ npm install
		$ IDE --headless & npm start

Take note of the pid ID to stop server process. Once your server is up feel free to use `npm start` directly.

### IE Fix
If build commands keep returning the same message, push the `F12` key down to open developer tools and try building again.
	
## Access

Visit [localhost:8884/index](http://localhost:8884/index). Access the IDE from any device on your network as well...

## Bug reports & questions :
Please create a new issue on Github to report a bug.

## Wiki : How to use the strukture

Visit https://github.com/thestrukture/IDE/wiki

## Misc info
How to install GoS incase the built-in installer fails.

- Install GoS [CLI](http://gophersauce.com). ( `$ go get github.com/cheikhshift/gos` )
- Install `GoS` dependencies : `$ gos deps`

### Extending
Automate your work flow with just lines of Javascript. Read the Guide [here](https://github.com/thestrukture/SpringMenu). 

## Contributions
Improvements to the codebase and pull requests are encouraged.

### Teams and small businesses
Get `IDE` setup for your business, with staff training. Learn more at [gophersauce.com](https://gophersauce.com).

###### Reporting
As a human I can't be everywhere, please help me find problems or unexpected behavior with this piece of software.

## Screenshots

![screenshot](tests/5500.png)
![screenshot](tests/5501.png)
![screenshot](tests/1newsc.png)
![screenshot](tests/2newsc.png)
![screenshot](tests/3newsc.png)
![screenshot](tests/4newsc.png)
![screenshot](tests/5newsc.png)
