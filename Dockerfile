#Base Image is golang
FROM golang:latest

# Set working directory to /go
WORKDIR /go/src/redfishmicro

#Copy go files
COPY Handlers.go Main.go Routes.go ./

EXPOSE 8090

RUN go get -u github.com/gorilla/mux

RUN go install

CMD /go/bin/redfishmicro
