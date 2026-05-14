# syntax=docker/dockerfile:1.7
FROM golang:1.23-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" -o /out/memorie ./cmd/memorie

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" -o /out/migrate ./cmd/migrate

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /out/memorie /memorie
COPY --from=builder /out/migrate /migrate

EXPOSE 8090
USER nonroot:nonroot

ENTRYPOINT ["/memorie"]
