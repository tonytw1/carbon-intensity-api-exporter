FROM golang:1.11

WORKDIR /go/src/app
COPY . .

CMD ["carbon-intensity-api-exporter"]

