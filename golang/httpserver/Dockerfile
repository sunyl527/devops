FROM golang:1.17 AS build
WORKDIR /app/
COPY main.go /app/main.go
RUN apt-get update &&  apt-get install dos2unix -y && dos2unix /app/main.go && GOOS=linux CGO_ENABLED=0 go build -o httpserver main.go

FROM busybox
COPY --from=build /app/httpserver /app/httpserver 
EXPOSE 80
WORKDIR /app/
CMD ["./httpserver"]
