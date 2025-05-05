FROM golang:1.24-alpine
WORKDIR /dbaas-api

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/api ./cmd/api 

CMD ["/dbaas-api/bin/api"]
EXPOSE 8080