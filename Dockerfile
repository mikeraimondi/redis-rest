FROM golang:1.10.3 as builder
WORKDIR /go/src/github.com/mikeraimondi/redis-rest
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /home/scratchuser
COPY --from=builder /go/src/github.com/mikeraimondi/redis-rest/app .

USER 10001
ENTRYPOINT [ "./app" ]
EXPOSE 8080
