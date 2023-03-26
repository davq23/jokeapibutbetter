# Set up base image
FROM golang:1.20.2-alpine as build-backend

RUN mkdir /app

ADD . /app

# Set up workdir
WORKDIR /app

# Copy and download dependencies
COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go build -o ./main ./main.go

# Build frontend on the cloud
FROM node:16.17.0-alpine as build-frontend
# Frontend assets path
ARG FRONTEND_ASSETS_URL=${FRONTEND_ASSETS_URL}
# API URL
ARG JOKEAPI_URL=${JOKEAPI_URL}

RUN if [[ -z "$FRONTEND_ASSETS_URL" ]] ; then export FRONTEND_ASSETS_URL=${FRONTEND_ASSETS_URL} ; fi
RUN if [[ -z "$JOKEAPI_URL" ]] ; then export VITE_JOKEAPI_URL=${JOKEAPI_URL} ; fi

RUN mkdir /app

ADD . /app

RUN mkdir /app/dist

# Set up workdir
WORKDIR /app

COPY vue-frontend/package*.json ./

RUN npm install

COPY vue-frontend/ .

RUN npm run build

# Run the API
FROM alpine

RUN mkdir /app
RUN mkdir /frontend
RUN apk update; apk add openssh
RUN addgroup -g 1001 appgroup && adduser -S -u 1001 -G appgroup appuser

COPY --from=build-frontend /app/dist /frontend
COPY --from=build-backend /app/main /main
COPY app_execute.sh /app_execute.sh

RUN chown -R appuser:appgroup /main
RUN chown -R appuser:appgroup /app_execute.sh

EXPOSE 8080

USER appuser

CMD ["/app_execute.sh"]





