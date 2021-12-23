FROM golang:1.17.5-buster

WORKDIR /go/src/app
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
RUN go get github.com/githubnemo/CompileDaemon

CMD [ "/entrypoint.sh" ]