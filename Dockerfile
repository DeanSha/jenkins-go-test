FROM golang:latest
#RUN mkdir /go
WORKDIR /go
ENV GOPATH /go
ENV GOBIN /go/bin
ADD . /go/
RUN go build -o $GOBIN/helloSvc hello
ENV PORT 8087
EXPOSE 8087
CMD ["/go/bin/helloSvc"]



