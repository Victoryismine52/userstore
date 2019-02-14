# pulling a lightweight version of golang (in Alpine Linux)
FROM golang:1.8-alpine
RUN apk --update add --no-cache git
# Copy the local package files to the container's workspace.
ADD . /go/src/userstore
WORKDIR /go/src/userstore
# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get userstore
# Run the command by default when the container starts.
ENTRYPOINT ["/go/bin/userstore"]
# Document that the service listens on port 9000.
EXPOSE 9000