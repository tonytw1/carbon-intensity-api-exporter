package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	currentIntensity = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "actual_intensity",
		Help: "Actual intensity ",
	})
)

func fetch() {
	go func() {
		for {
			api_url := "https://api.carbonintensity.org.uk"
			current_intensity_url := api_url + "/intensity"

			req, err := http.NewRequest(http.MethodGet, current_intensity_url, nil)
			if err != nil {
				log.Fatal(err)
			}

			client := http.Client{
				Timeout: time.Second * 5,
			}

			res, getErr := client.Do(req)
			if getErr != nil {
				log.Fatal(getErr)
			}

			body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
				log.Fatal(readErr)
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

			current := IntensityData{}
			err = json.Unmarshal(body, &current)
			if err != nil {
				log.Fatal(err)
			}

			data := current.Data[0]
			currentIntensity.Set(float64(data.Intensity.Actual))

			time.Sleep(60 * time.Second)
		}
	}()
}

func main() {
	fetch()

	minimalRegistry := prometheus.NewRegistry()
	minimalRegistry.MustRegister(currentIntensity)
	handler := promhttp.HandlerFor(minimalRegistry, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)
	http.ListenAndServe(":8080", nil)
}
