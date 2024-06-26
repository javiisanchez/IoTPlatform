# # A sample microservice in Go packaged into a container image.
# FROM golang:1.22.3

# # Set destination for COPY
# WORKDIR /app

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

# # Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-iotplatform-docker

# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 8080

# Run
#CMD ["/golang-iotplatform-docker"]


FROM golang:1.22.3

WORKDIR /go/src/myapp

COPY . .

#RUN go build -o /go/bin/myapp cmd/main.go

