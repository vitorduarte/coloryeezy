# Build the program
FROM golang:alpine AS builder

# Install git.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/vitorduarte/coloryeezy/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build


# Build a small image
FROM alpine:latest
COPY --from=builder /go/src/vitorduarte/coloryeezy/coloryeezy /coloryeezy/coloryeezy
COPY --from=builder /go/src/vitorduarte/coloryeezy/fonts /coloryeezy/fonts

WORKDIR /coloryeezy/

# Run the hello binary.
ENTRYPOINT ["/coloryeezy/coloryeezy"]
