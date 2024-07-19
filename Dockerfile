FROM docker.arvancloud.ir/golang:latest

WORKDIR /app

COPY app/go.mod app/go.sum ./

RUN go mod download

COPY app ./

RUN go build -o out

EXPOSE 8080

CMD ["/out"]

