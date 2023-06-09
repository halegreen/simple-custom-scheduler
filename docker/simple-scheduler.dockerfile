FROM golang:1.17 as builder
WORKDIR /go/src/github.com/halegreen/simple-custom-scheduler

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd/ cmd/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o koord-descheduler cmd/koord-descheduler/main.go

FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /go/src/github.com/halegreen/simple-custom-scheduler .
ENTRYPOINT ["/simple-scheduler"]
