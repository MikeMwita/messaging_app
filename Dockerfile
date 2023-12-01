FROM golang:1.21rc2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/api/main.go


# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
COPY start.sh .
COPY wait-for.sh /app/

ENV DB_CONNECTION_STRING ""

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]