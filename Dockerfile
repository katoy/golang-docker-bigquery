# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/golang/example/outyet
ADD client.json /go

RUN go get -u golang.org/x/net/context
RUN go get -u golang.org/x/oauth2/google
RUN go get -u google.golang.org/api/bigquery/v2
# RUN go get -u google.golang.org/api/storage/...
# RUN go get -u google.golang.org/api/bigquery/...

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/golang/example/outyet

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/outyet
# Document that the service listens on port 8080.
EXPOSE 8080
