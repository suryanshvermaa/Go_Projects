FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main .

FROM golang:1.24-alpine AS runner

WORKDIR /app
COPY --from=builder /build/main /app/

RUN addgroup --system --gid 1001 go
RUN adduser --system --uid 1001 go
USER go
EXPOSE 3000

CMD ["./main"]