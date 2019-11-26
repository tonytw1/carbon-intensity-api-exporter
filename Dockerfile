FROM golang:1.11

WORKDIR /go/src/app
COPY . .

CMD ["/go/src/app/carbon-intensity-api-exporter"]

