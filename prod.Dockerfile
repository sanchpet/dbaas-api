# Build environment
# -----------------
FROM golang:1.24-alpine AS build-env
WORKDIR /dbaas-api

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api


# Deployment environment
# ----------------------
FROM alpine

COPY --from=build-env /dbaas-api/bin/api /dbaas-api/

EXPOSE 8080
CMD ["/dbaas-api/api"]