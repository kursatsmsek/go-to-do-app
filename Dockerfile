FROM golang

WORKDIR /src
EXPOSE 8080
ADD . .
RUN go mod download && go build
ENTRYPOINT ["./todo-app"]