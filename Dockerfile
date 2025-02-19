# Build stage
FROM golang:1.23.4-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
# 위에서 생성한 main 파일을 실행하기 위한 5MB인 Alpine Linux
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]