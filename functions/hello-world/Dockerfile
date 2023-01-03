 # Use the official Golang image to create a build artifact.
 # This is based on Debian and sets the GOPATH to /go.
 # https://hub.docker.com/_/golang
 FROM golang:1.19 as builder

 # Copy local code to the container image.
 WORKDIR /app

 # Copy local code to the container image.
 COPY . ./

 RUN make build

 # Use a Docker multi-stage build to create a lean production image.
 # https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
 FROM alpine:3
 RUN apk add --no-cache ca-certificates

 # Copy the binary to the production image from the builder stage.
 COPY --from=builder /app/app /app

 # Run the web service on container startup.
 CMD ["/app"]
