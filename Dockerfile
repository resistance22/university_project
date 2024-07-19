FROM docker.arvancloud.ir/golang:latest as builder

WORKDIR /app

COPY app/go.mod app/go.sum ./

RUN go mod download

COPY app ./

RUN go build -o out

FROM docker.arvancloud.ir/alpine:latest as runner 

WORKDIR /app

RUN apk add libc6-compat

COPY dev.env /

COPY --from=builder /app/out . 

EXPOSE 8080

CMD ["./out"]

