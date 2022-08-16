# Builder
FROM golang:1.17-alpine AS builder
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
ENV PROJECT miao-sha
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN chmod 777 /app/run.sh
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o miao-sha
EXPOSE 8081
ENTRYPOINT ["/app/run.sh"]
