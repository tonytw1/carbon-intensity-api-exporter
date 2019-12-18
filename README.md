## Carbon Intensity API exporter

Exports [UK Carbon Intensity API](https://carbonintensity.org.uk/) readings in Prometheus format.


### Build and run locally

```
https://github.com/tonytw1/carbon-intensity-api-exporter.git

go get -d -v ./...
go install -v  ./...
go build ./...

./carbon-intensity-api-exporter
```

### Run with Docker

```
docker pull eu.gcr.io/eelpie-public/carbon-intensity-api-exporter
docker run -p 8080:8080 eu.gcr.io/eelpie-public/carbon-intensity-api-exporter
```


### Sample output

```
http://localhost:8080
```
