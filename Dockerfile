# syntax=docker/dockerfile:1.7
FROM golang:1.23-alpine AS builder

WORKDIR /src
COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags="-s -w" -o /out/memorie ./cmd/memorie

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /out/memorie /memorie

EXPOSE 8090
USER nonroot:nonroot

ENTRYPOINT ["/memorie"]
