FROM golang:1.26-alpine AS build

WORKDIR /workspace

COPY . .

RUN go mod tidy

RUN go build -o backend cmd/main.go

RUN chmod +x backend


FROM alpine:latest

WORKDIR /app

COPY --from=build /workspace/backend /app

EXPOSE 8888

ENTRYPOINT ["/app/backend"]  