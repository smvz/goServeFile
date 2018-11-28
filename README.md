# goServeFile

Simple tool written in Go to serve one file over HTTP. Useful for transferring individual, large files over the Internet.

## Installation

Install the Go programming language compiler for your platform, e.g. for Ubuntu 18.04:
`sudo apt-get install golang`

Compile the source:
`go build goServeFile.go`

## Example Usage

`echo "test" > testfile.txt && ./goServeFile testfile.txt`

```
$ curl localhost:8080
<a href="/testfile.txt">Found</a>.
$ curl -L localhost:8080
test
```

Ctrl+c to stop listening and terminate the program

## TODO

* Support resuming downloads
* Support choosing directories, presenting a directory listing to clients by default
* Add option to choose a directory, archiving the contents and providing a single download of that archive
