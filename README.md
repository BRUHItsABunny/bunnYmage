# bunnYmage
A wrapper written for Golang image processing, sparked for  my need for a half decent wrapper but not being able to find one.


##Introduction
Introduction to what bunnYmage is and why it exists.

#####Why is bunnYmage?
I had the need of converting images and turning a list of them into a GIF, upon researching how to do that in pure Golang I came to the conclusion that it was overly complex for my taste.

#####What is bunnYmage?
It is a pure Golang solution for image processing, with the key features being file converting and GIF creation.
I want to add drawing and other features too but those are most definitely not on my prioritized TODO list.

#####When is bunnYmage?
This is my road map for bunnYmage.
* File convertion and global image classes (without writing to disk preferably) Done: [05/05/2019]
* GIF creation Done: [05/05/2019]
* Image editing
    * Cropping images
    * Resizing images
    * Drawing on images
    * Overlaying images

##Installation
Just run the following:

```go get github.com/tmjvonboss/bunnYmage```

That should be enough to get started.

##Examples
Examples can be found here.

##Known issues
A few issues I spotted myself that I have yet to fix.
* Quality drop on GIF generation uncontrolled.
##Support
Help is always welcome.

##Usage
I will add a license, pretty much allows you to use it in non-commercial projects freely, would preferably like to know what you use it for so I could add it to a list of projects this was used in.


##Tested on
Golang version

```go version go1.12.4 windows/amd64```

Golang env

```
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\BUNNY\AppData\Local\go-build
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOOS=windows
set GOPATH=C:\Users\BUNNY\go
set GOPROXY=
set GORACE=
set GOROOT=C:\Go
set GOTMPDIR=
set GOTOOLDIR=C:\Go\pkg\tool\windows_amd64
set GCCGO=gccgo
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\BUNNY\AppData\Local\Temp\go-build066875722=/tmp/go-build -gno-record-gcc-switches
```
