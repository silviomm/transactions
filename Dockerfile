FROM ubuntu:latest
LABEL authors="silviomm"

# Build
FROM golang:1.16-buster as builder
WORKDIR /app
COPY . .
RUN go build -o pismo-api pismo-challenge

# Run
FROM golang:1.16-buster
WORKDIR /app
COPY --from=builder /app/.env .
COPY --from=builder /app/pismo-api .
EXPOSE 3000
CMD ["./pismo-api"]
