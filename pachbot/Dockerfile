FROM golang:1.9.2 as builder
WORKDIR  /go/src/pachbot/
COPY . /go/src/pachbot/
#RUN go test ./... -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run .

FROM alpine:3.8
RUN apk --update add ca-certificates && apk --update add iptables
WORKDIR /root/
COPY --from=builder /go/src/pachbot/run .
CMD ["./run"]