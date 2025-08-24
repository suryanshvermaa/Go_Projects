FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main

FROM alpine:3.14 AS runner

WORKDIR /app

COPY --from=builder /build/main /app/main

RUN addgroup --system --gid 1001 golang
RUN adduser --system --uid 1001 golang

USER golang
EXPOSE 3000
ENV PORT=3000

CMD [ "/app/main" ]

