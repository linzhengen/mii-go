FROM golang:1.24 as build

WORKDIR /src
COPY go.mod ./
COPY go.sum ./
COPY ./ ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app /src/cmd/app

FROM gcr.io/distroless/base:debug-nonroot
WORKDIR /

COPY --from=build /app /app

ENTRYPOINT ["/app", "rest"]