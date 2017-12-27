# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8

# Copy the local package files to the container's workspace.
ADD . /go/src/intlogs

# setup dependencies
WORKDIR /go/src/intlogs
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

RUN go install intlogs


# Run the command by default when the container starts.
ENTRYPOINT /go/bin/intlogs /go/src/intlogs/config.yml

# Document that the service listens on port 8080.
EXPOSE 8000