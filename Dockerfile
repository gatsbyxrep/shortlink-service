FROM golang:latest

RUN apt-get update && apt-get install -y postgresql-client

WORKDIR /app

COPY . .

RUN go build -o app .

RUN go get github.com/lib/pq

COPY init.sql /docker-entrypoint-initdb.d/
RUN chmod +x /docker-entrypoint-initdb.d/init.sql

CMD ["./app"]