# Set up base image
FROM golang:1.19-alpine as build-backend

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

FROM node:16.17.0-alpine as build-frontend

RUN mkdir /app

ADD . /app

RUN mkdir /app/dist

# Set up workdir
WORKDIR /app

COPY vue-frontend/package*.json ./

RUN npm install

COPY vue-frontend/ .

RUN npm run build

# Actually running the app
FROM alpine

RUN mkdir /dist

#RUN addgroup -g 1001 appgroup && adduser -S -u 1001 -G appgroup appuser


#RUN chown -R appuser:appgroup /app /dist
#RUN chmod 755 /app /app/dist

COPY --from=build-backend /app/main /main
COPY --from=build-frontend /app/dist/ /dist

EXPOSE 8080

#USER appuser

CMD ["/main"]





