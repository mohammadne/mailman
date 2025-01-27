FROM golang:1.17 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /bin/app ./cmd/root.go

FROM alpine:latest

RUN apk add --no-cache libc6-compat 

WORKDIR /bin/

COPY --from=builder /bin/app .

ARG COMMAND

LABEL org.opencontainers.image.source="https://github.com/mohammadne/middleman/${COMMAND}"

ENTRYPOINT ["/bin/app"]

CMD [${COMMAND}, "--env=prod"]
