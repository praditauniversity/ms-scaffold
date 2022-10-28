FROM golang:alpine
RUN mkdir /scaffold
WORKDIR /scaffold

ADD ./go.mod .
ADD ./go.sum .
RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN mkdir /scaffold/bin

ADD . .

EXPOSE 8080
ENTRYPOINT CompileDaemon -log-prefix=false -directory="." -build="go build -o bin/scaffold" -command="./bin/scaffold"