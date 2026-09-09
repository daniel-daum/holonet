# ---- build ----
FROM golang:1.27.1-alpine AS build

WORKDIR /src

# dependencies first: editing Go source won't re-run this step
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o /out/holonet ./cmd/holonet

# ---- runtime ----
FROM alpine:3.24

# ca-certificates for outbound TLS, curl for poking at the service from inside
RUN apk add --no-cache ca-certificates curl \
    && adduser -D -u 65532 holonet

COPY --from=build /out/holonet /usr/local/bin/holonet

USER holonet
EXPOSE 7432

ENTRYPOINT ["holonet"]
