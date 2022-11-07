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


# Build ratings service
RUN go build -o ./ratings ./services/ratings/main.go

# Build main service 

FROM alpine

COPY --from=build /app/ratings /ratings
COPY --from=build /app/db.env /db.env

EXPOSE 8056

CMD ["/ratings"]