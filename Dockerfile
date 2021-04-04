FROM golang:1.15.7-alpine

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

RUN apk update && apk add git
COPY . ${ROOT}

RUN go mod download

EXPOSE 8080

RUN go get -u github.com/codegangsta/gin
#RUN go get -u github.com/cosmtrek/air
#CMD ["air", "-c", ".air.toml"]