FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /app

COPY . .

# Build
RUN go build -o app_hris_server

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD ./app_hris_server
