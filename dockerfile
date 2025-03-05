#FROM golang:alpine
#
#WORKDIR /app
#COPY go.mod ./
#RUN go mod download
#COPY . .
#RUN go build -o /app/unified-auth-system ./cmd/uasbreezy/main.go
#EXPOSE 8008
#CMD ["./unified-auth-system"]

FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/unified-auth-system ./cmd/uasbreezy/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/unified-auth-system /app/unified-auth-system
COPY config ./config
EXPOSE 8008
CMD ["./unified-auth-system"]