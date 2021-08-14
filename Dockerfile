FROM golang:alpine3.14 as builder


WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
COPY main.go .
COPY ./code ./code
COPY main.go .

RUN go mod download
RUN go build -o linuxExe main.go



FROM alpine:3.13


WORKDIR /opt/myapp/
COPY --from=builder /go/src/app/linuxExe .

COPY ./templates ./templates

RUN chmod +x linuxExe

EXPOSE 80

ENTRYPOINT [ "./linuxExe" ]