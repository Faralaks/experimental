FROM golang:1.15.2


WORKDIR /go/src/app
COPY main.go .

RUN go get ./..
RUN go build -o go_tst main.go
RUN chmod +x go_tst

EXPOSE 80


CMD [ "./go_tst" ]