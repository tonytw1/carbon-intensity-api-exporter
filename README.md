## Carbon Intensity API exporter

Exports [UK Carbon Intensity API](https://carbonintensity.org.uk/) readings in Prometheus format.


### Build and run locally

```
git clone https://github.com/tonytw1/carbon-intensity-api-exporter.git

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

### Scrape public target

```
- job_name: 'carbon-intensity'
  static_configs:
  - targets: ['carbon-intensity.eelpieconsulting.co.uk']
  scrape_interval: 60s
```

### Sample output

```
http://localhost:8080

# HELP actual_intensity Actual intensity 
# TYPE actual_intensity gauge
actual_intensity 206
# HELP forecast_intensity Forecast intensity 
# TYPE forecast_intensity gauge
forecast_intensity 204
```
