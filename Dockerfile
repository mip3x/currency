ARG GOLANG_VERSION="1.22"
ARG ALPINE_VERSION="3.21"

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ARG LDFLAGS="-s -w"

RUN go build -ldflags="${LDFLAGS}" -o main main.go


FROM alpine:${ALPINE_VERSION}

WORKDIR /app

COPY --from=builder /src/main .

ARG APP_PORT=8080
EXPOSE ${APP_PORT}

RUN adduser -D appuser && chown -R appuser /app
USER appuser

CMD ["./main"]
