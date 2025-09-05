FROM golang:1.25.1-alpine3.21 AS builder

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod tidy

COPY . .

RUN go build -o main

FROM alpine:3.22.1 AS runner

WORKDIR /app
COPY --from=builder /build/main /app/main

RUN addgroup --system --gid 1001 golang
RUN adduser --system --uid 1001 golang

USER golang
EXPOSE 3000
ENV PORT=3000

CMD [ "/app/main" ]
