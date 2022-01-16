FROM golang:1.17 AS build
WORKDIR /app/
COPY main.go /app/main.go
RUN go build -o httpserver main.go

FROM busybox
COPY --from=build /app/httpserver /app/httpserver 
EXPOSE 80
WORKDIR /app/
CMD ["./httpserver"]
