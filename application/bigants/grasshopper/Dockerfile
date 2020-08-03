FROM golang:1.13.4-buster as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
ADD . ./
RUN GOOS=linux go build -a -installsuffix cgo -o migrate github.com/allscape/bigants/grasshopper/cmd/migrate
RUN GOOS=linux go build -a -installsuffix cgo -o server github.com/allscape/bigants/grasshopper/cmd/server

FROM debian:bullseye-20190708
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /root/
COPY --from=builder /app/server ./
RUN mkdir -p app/db/
COPY --from=builder /app/app/db/migrations ./app/db/migrations
CMD "/root/server"
