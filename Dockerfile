# Build image
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /main

# Runtime image 
FROM scratch
COPY --from=builder /main /main
EXPOSE 8080/tcp
ENTRYPOINT ["/main"]
