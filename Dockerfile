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
ARG FRONTEND_ASSETS_URL
# API URL
ARG JOKEAPI_URL

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

# Upload Vue frontend through FTP
FROM php:8.1.17RC1-cli-alpine3.16 as upload-frontend-ftp
ARG FTP_HOST
ARG FTP_PORT
ARG FTP_USERNAME
ARG FTP_PASSWORD
ARG FTP_TIMEOUT

RUN if [[ -z "$FTP_HOST" ]] ; then export FTP_HOST=${FTP_HOST} ; fi
RUN if [[ -z "$FTP_PASSWORD" ]] ; then export FTP_PASSWORD=${FTP_PASSWORD} ; fi
RUN if [[ -z "$FTP_USERNAME" ]] ; then export FTP_USERNAME=${FTP_USERNAME} ; fi
RUN if [[ -z "$FTP_PORT" ]] ; then export FTP_PORT=${FTP_PORT} ; fi

RUN mkdir /upload
RUN mkdir /upload/dist

COPY upload-frontend-ftp/ /upload

COPY --from=build-frontend /app/dist/ /upload/dist

RUN cd upload/; php ftp_upload.php

RUN touch finish.txt

# Run the API
FROM alpine

RUN mkdir /app

RUN addgroup -g 1001 appgroup && adduser -S -u 1001 -G appgroup appuser

COPY --from=upload-frontend-ftp /finish.txt /finish.txt
COPY --from=build-backend /app/main /main

RUN chown -R appuser:appgroup /main

EXPOSE 8080

USER appuser

CMD ["/main"]





