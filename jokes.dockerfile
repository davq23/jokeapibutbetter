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

# Build joke service
RUN go build -o ./jokes ./services/jokes/main.go

# Execute build
FROM alpine

COPY --from=build /app/jokes /jokes
COPY --from=build /app/db.env /db.env

EXPOSE 8055

CMD ["/jokes"]