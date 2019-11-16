FROM golang:1.12.6-alpine3.10 AS build

WORKDIR /src/
COPY go-server.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo

FROM scratch
COPY --from=build /bin/demo /bin/demo
ENTRYPOINT ["/bin/demo"]
