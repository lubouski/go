## How-to
We using Docker multistage build, and `goland-alpine` image to build binary and `stretch` to run it.

In this example is easy to understand how to use built in `go` server and `os/exec` command to operating system

### To run and build, just follow the instruction bellow.
```
$ sudo docker build -t my-web-server .
```
```
$ sudo docker run -d -p 9999:8080 my-web-server
```
```
$ curl http://localhost:9999/
```
```
Whoa Go is neat! today 2019-11-16
```
## Alternatively binary could be run via go runtime
```
$ go run go-server.go 
```

