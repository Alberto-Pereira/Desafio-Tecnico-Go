# syntax=docker/dockerfile:1

# Golang version
FROM golang:1.18

# Image dir
WORKDIR /app

# Copies the files that contain dependencies
COPY go.mod ./
COPY go.sum ./

# Download the dependencies
RUN go mod download

# Copies all files from api
COPY . .

# API Build
RUN go build -o /docker-desafio-tecnico-go

# User port : Container port
EXPOSE 8080:8080

# Executes
CMD [ "/docker-desafio-tecnico-go" ]