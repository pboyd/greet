FROM golang:1.26-bookworm AS builder
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o /hello .

FROM scratch
COPY --from=builder /hello /hello
ENTRYPOINT ["/hello"]
