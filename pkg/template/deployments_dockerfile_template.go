package template

type dockerfileTemplate struct {
	Module  string
	Project string
}

func (d *dockerfileTemplate) Text() []byte {
	return []byte(`# build stage
FROM golang:alpine AS builder

ENV GO111MODULE=on

# install git.
RUN apk update && apk add --no-cache git

RUN mkdir -p /go/src/` + d.Module + `
WORKDIR /go/src/` + d.Module + `
COPY . .

# Download all dependencies
RUN go mod download

# With go ≥ 1.10
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/` + d.Project + ` cmd/` + d.Project + `/main.go

# small image
FROM alpine:3.7

WORKDIR /app
COPY --from=builder /go/src/` + d.Module + `/configuration/production.yml /app/configuration/production.yml
COPY --from=builder /go/src/` + d.Module + `/policy /app/policy
COPY --from=builder /go/bin/` + d.Project + ` .

ENV TZ=Asia/Bangkok
RUN echo "Asia/Bangkok" > /etc/timezone

# run binary.
ENTRYPOINT ["/app/` + d.Project + `", "-env", "production"]`)
}

func DeploymentsDockerfileTemplate(module string, project string) Template {
	return &dockerfileTemplate{
		Module:  module,
		Project: project,
	}
}
