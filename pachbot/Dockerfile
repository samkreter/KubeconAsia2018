FROM golang:1.9.2 as builder

RUN apt update
RUN apt install -y curl

# Install pachctl 
RUN curl -o /tmp/pachctl.tar.gz -L https://github.com/pachyderm/pachyderm/releases/download/v1.7.10/pachctl_1.7.10_linux_amd64.tar.gz && tar -xvf /tmp/pachctl.tar.gz -C /tmp && cp /tmp/pachctl_1.7.10_linux_amd64/pachctl /usr/local/bin


WORKDIR  /go/src/github.com/samkreter/KubeconAsia2018/pachbot/
COPY . /go/src/github.com/samkreter/KubeconAsia2018/pachbot/
#RUN go test ./... -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run .
CMD ["./run"]