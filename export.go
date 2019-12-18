package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const apiUrl = "https://api.carbonintensity.org.uk"
const currentIntensityUrl = apiUrl + "/intensity"
const pollingInterval = 5 * time.Minute

var (
	currentIntensity = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "actual_intensity",
		Help: "Actual intensity ",
	})
	forecastIntensity = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "forecast_intensity",
		Help: "Forecast intensity ",
	})
)

var client = http.Client{
	Timeout: time.Second * 10,
}

type Intensity struct {
	Forecast int
	Actual   int
	Index    string
}

type Data struct {
	From      string
	To        string
	Intensity Intensity
}

type IntensityData struct {
	Data []Data
}

func poll() {
	go func() {
		for {
			err := update()
			if err != nil {
				log.Print("Error while updating: ", err)
			}
			time.Sleep(pollingInterval)
		}
	}()
}

func update() error {
	body, err := fetch(currentIntensityUrl)
	if err != nil {
		return err
	}
	current := IntensityData{}
	err = json.Unmarshal(body, &current)
	if err != nil {
		return err
	}

	if len(current.Data) > 0 {
		data := current.Data[0]
		currentIntensity.Set(float64(data.Intensity.Actual))
		forecastIntensity.Set(float64(data.Intensity.Forecast))
	}
	return nil
}

func fetch(url string) ([]byte, error) {
	log.Print("Fetching from " + url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}
	return body, err
}

func main() {
	log.SetOutput(os.Stdout)

	minimalRegistry := prometheus.NewRegistry()
	minimalRegistry.MustRegister(currentIntensity, forecastIntensity)

	poll()

	handler := promhttp.HandlerFor(minimalRegistry, promhttp.HandlerOpts{})
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
