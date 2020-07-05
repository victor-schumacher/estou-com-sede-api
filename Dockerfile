FROM golang:latest

WORKDIR $GOPATH/src/github.com/victor-schumacher/estou-com-sede-api

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...
##Tes    
EXPOSE $PORT

CMD ["estou-com-sede-api"]