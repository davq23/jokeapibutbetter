# Set up base image
FROM golang:1.19-alpine as build

RUN mkdir /app

COPY ./launch.sh /app/launch.sh

ADD . /app

# Set up workdir
WORKDIR /app

# Copy and download dependencies
COPY go.mod ./
COPY go.sum ./
COPY launch.sh ./

RUN go mod download

# Build config service
# RUN go build -o ./config /services/config/main.go

# Build joke service
RUN go build -o ./jokes ./services/jokes/main.go

# Build ratings service
RUN go build -o ./ratings ./services/ratings/main.go

# Build users service
RUN go build -o ./users ./services/users/main.go

# Build main service 
# RUN go build -o ./main /services/frontapi/main.go

FROM alpine

COPY --from=build /app/users /users
COPY --from=build /app/jokes /jokes
COPY --from=build /app/ratings /ratings
COPY --from=build /app/launch.sh /launch.sh
COPY --from=build /app/db.env /db.env

# Expose public port
EXPOSE 8056
EXPOSE 8055

#"/cmd/main" ]
CMD ["/bin/sh","./launch.sh"]
