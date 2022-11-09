# Set up base image
FROM golang:1.19-alpine as build

RUN mkdir /app

ADD . /app

# Set up workdir
WORKDIR /app

# Copy and download dependencies
COPY go.mod ./
COPY go.sum ./
COPY local.env ./

RUN go mod download


# Build ratings service
RUN go build -o ./config ./services/config/main.go

# Build main service 

FROM alpine

COPY --from=build /app/config /config
COPY --from=build /app/local.env /local.env

EXPOSE 8054

CMD ["/config"]

HEALTHCHECK --interval=5s --timeout=3s \
  CMD wget http://localhost:8054/config?format=json -o config.json & rm config.json || exit 1