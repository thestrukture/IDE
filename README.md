# Strukture IDE
[![Build Status](https://travis-ci.org/thestrukture/IDE.svg?branch=master)](https://travis-ci.org/thestrukture/IDE)
[![GoDoc](https://godoc.org/github.com/thestrukture/IDE/api?status.svg)](https://godoc.org/github.com/thestrukture/IDE/api)
[![Test Coverage](https://api.codeclimate.com/v1/badges/d46b0bfb51e827632710/test_coverage)](https://codeclimate.com/github/thestrukture/IDE/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/d46b0bfb51e827632710/maintainability)](https://codeclimate.com/github/thestrukture/IDE/maintainability)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fthestrukture%2FIDE.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fthestrukture%2FIDE?ref=badge_shield)

## About IDE
Just a Go IDE. It features :
- Autocomplete.
- Syntax correction.
- Interactive terminal via web socket.
- Line tags in relation to failed build logs.
- Project build scripts.
- View web application output.
- Build docker images (Must have docker running on host) . 
- Breakpoints and debugging with Delve.
- Regex directory search.
## Requirements
- Go v1.15+.
- Git. Git present as a command on your system.

## First Launch

If the server launch hangs on startup, close it and install the additional requirements manually. Prior to running the commands, set your GOPATH to `$home/workspace`. You can do this on Windows with `set GOPATH=%USERPROFILE%\workspace`.
Run the following command : 

	go get github.com/mdempsky/gocode

To add debug support, you must install delve. You can find the guide here. (Don't worry it is quick) [Install Delve](https://github.com/go-delve/delve/tree/master/Documentation/installation)

## Install

		$ go install github.com/thestrukture/IDE@latest

#### How to run

		$ IDE
		
#### Run as a server

		$ IDE --headless

## Access

Visit [localhost:8884/index](http://localhost:8884/index). Access the IDE from any device on your network as well...

## Bug reports & questions :
Please create a new issue on Github to report a bug.

## How to use the strukture

Get training for $299.99 [https://thestrukture.gumroad.com/l/strukture_live_training](https://thestrukture.gumroad.com/l/strukture_live_training)

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fthestrukture%2FIDE.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fthestrukture%2FIDE?ref=badge_large)
