FROM  golang:1.23.1-bullseye AS builder

WORKDIR /app

EXPOSE 8080

RUN --mount=source=go.mod,target=go.mod \
    --mount=source=go.sum,target=go.sum \
    go mod download

RUN --mount=source=.,target=. \
    go build -o /go/bin/main .


FROM gcr.io/distroless/cc:latest
EXPOSE 8080

COPY --from=builder /go/bin/main /go/bin/main

CMD ["/go/bin/main"]