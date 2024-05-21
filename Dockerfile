FROM golang AS build
WORKDIR /build
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod/cache \
  go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 go build -o rest-api cmd/rest-api/main.go

FROM scratch
COPY migrations migrations
COPY --from=build /build/rest-api .
CMD ["./rest-api"]
