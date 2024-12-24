FROM golang:1.23.4

# Set destination for COPY
WORKDIR /app

# Download Go modules
# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
#COPY *.go ./
COPY . ./

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
# COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /fplbuddy


# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/fplbuddy"]