# build stage
FROM golang:alpine AS builder

ENV GO111MODULE=on

# install git.
RUN apk update && apk add --no-cache git

RUN mkdir -p /go/src/github.com/prongbang/mvp/test_test
WORKDIR /go/src/github.com/prongbang/mvp/test_test
COPY . .

# Using go mod with go 1.11
RUN go mod vendor

# With go ≥ 1.10
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/test_test cmd/test_test/main.go

# small image
FROM alpine:3.7

WORKDIR /app
COPY --from=builder /go/src/github.com/prongbang/mvp/test_test/configuration/production.yml /app/configuration/production.yml
COPY --from=builder /go/src/github.com/prongbang/mvp/test_test/policy /app/policy
COPY --from=builder /go/bin/test_test .

ENV TZ=Asia/Bangkok
RUN echo "Asia/Bangkok" > /etc/timezone

# run binary.
ENTRYPOINT ["/app/test_test", "-env", "production"]