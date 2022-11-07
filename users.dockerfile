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

# Build main service 
# RUN go build -o ./main /services/frontapi/main.go

FROM alpine

COPY --from=build /app/users /users
COPY --from=build /app/db.env /db.env

# Expose public port
EXPOSE 8080

#"/cmd/main" ]
CMD ["/users"]
