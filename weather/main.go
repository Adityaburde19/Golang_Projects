package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/",
			func(w http.ResponseWriter, r *http.Request){
				city := strings.SplitN(URL.Path, "/", 3)[2]
			}
		)

	http.ListenAndServe("localhost:8080", nil)

}

type apiConfigData struct {
	OpenWeatherMapAiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}
	return c, nil

}
