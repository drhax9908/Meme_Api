FROM golang:1.13.5

WORKDIR /go/src/github.com/drhax9908/Meme_Api

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD [ "go", "run", "." ]
