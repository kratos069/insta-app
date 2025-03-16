# Build Stage 
FROM golang:1.24.1-alpine3.21 AS builder
 WORKDIR /app
#  copy from root (where Dockerfile ran) project to current work dir (/app)
COPY . .
#  output binary file (main), o is output, current main file = main.go
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-386.tar.gz | tar xvz

# Run Stage
FROM alpine:3.21
WORKDIR /app
# /app/main is the path to the file we want to copy, "." is current work dir
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]