## How-to
We using Docker multistage build, and `goland-alpine` image to build binary and `stretch` to run it.

### To run and build, just follow the instruction bellow.
```
$ sudo docker build -t mytestbox .
```
```
$ sudo docker run -d -p 9999:8888 mytestbox
```
```
$ curl http://localhost:9999/
```
```
Hello, 世界
```
### That's all folks! 
