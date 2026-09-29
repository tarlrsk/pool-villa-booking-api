FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

# Minimal runtime image; includes CA certs needed for Neon TLS and the LINE API.
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/api /api
ENTRYPOINT ["/api"]
