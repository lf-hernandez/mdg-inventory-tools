# Stage 1: Build
FROM golang:1.21-alpine as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Stage 2: Run
FROM alpine:latest

WORKDIR /root/
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]
