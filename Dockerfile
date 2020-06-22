FROM golang

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN go get github.com/daflad/scrabble.nehpets.co.uk
WORKDIR /go/src/github.com/daflad/scrabble.nehpets.co.uk

ENTRYPOINT ./puller.sh & 2>&1 && revel run github.com/daflad/scrabble.nehpets.co.uk

EXPOSE 2099
