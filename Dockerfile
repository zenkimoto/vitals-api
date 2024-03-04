FROM golang:latest

RUN mkdir /go/src/work
WORKDIR /go/src/work

COPY . /go/src/work

RUN make build

EXPOSE 8080

CMD ["make", "run-build"]
