# Set up base image
FROM golang:1.19-alpine as build

RUN mkdir /app

ADD . /app

# Set up workdir
WORKDIR /app

# Copy and download dependencies
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# Build users service
RUN go build -o ./users ./services/users/main.go

FROM alpine

COPY --from=build /app/users /users

# Expose public port
EXPOSE 8080

#"/cmd/main" ]
CMD ["/users"]
