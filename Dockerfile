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

# Build config service
RUN go build -o ./main ./main.go

# Actually running the app
FROM alpine

COPY --from=build /app/main /main

EXPOSE 8080

CMD ["/main"]





