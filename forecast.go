package darksky

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// URL example:  "https://api.darksky.net/forecast/APIKEY/LATITUDE,LONGITUDE,TIME?units=ca&lang=en"
const (
	BASEURL = "https://api.darksky.net/forecast"
)

//Flags are the optional fields that can come back from the API such as weather stations
type Flags struct {
	DarkSkyUnavailable string   `json:"darksky-unavailable,omitempty"`
	DarkSkyStations    []string `json:"darksky-stations,omitempty"`
	DataPointStations  []string `json:"datapoint-stations,omitempty"`
	ISDStations        []string `json:"isds-stations,omitempty"`
	LAMPStations       []string `json:"lamp-stations,omitempty"`
	MADISStations      []string `json:"madis-stations,omitempty"`
	METARStations      []string `json:"metars-stations,omitempty"`
	METNOLicense       string   `json:"metnol-license,omitempty"`
	Sources            []string `json:"sources,omitempty"`
	Units              string   `json:"units,omitempty"`
}

// DataPoint is the catch all type that includes all of the fields a forecast could have
type DataPoint struct {
	Time                       int     `json:"time,omitempty"`
	Summary                    string  `json:"summary,omitempty"`
	Icon                       string  `json:"icon,omitempty"`
	SunriseTime                int     `json:"sunriseTime,omitempty"`
	SunsetTime                 int     `json:"sunsetTime,omitempty"`
	MoonPhase                  float64 `json:"moonPhase,omitempty"`
	PrecipIntensity            float64 `json:"precipIntensity,omitempty"`
	PrecipIntensityMax         float64 `json:"precipIntensityMax,omitempty"`
	PrecipIntensityMaxTime     int     `json:"precipIntensityMaxTime,omitempty,omitempty"`
	PrecipProbability          float64 `json:"precipProbability,omitempty"`
	PrecipAccumulation         float64 `json:"precipAccumulation,omitempty"`
	PrecipType                 string  `json:"precipType,omitempty,omitempty"`
	Temperature                float64 `json:"temperature,omitempty"`
	TemperatureMin             float64 `json:"temperatureMin,omitempty"`
	TemperatureMinTime         int     `json:"temperatureMinTime,omitempty"`
	TemperatureMax             float64 `json:"temperatureMax,omitempty"`
	TemperatureMaxTime         int     `json:"temperatureMaxTime,omitempty"`
	ApparentTemperature        float64 `json:"apparentTemperature,omitempty"`
	ApparentTemperatureMin     float64 `json:"apparentTemperatureMin,omitempty"`
	ApparentTemperatureMinTime int     `json:"apparentTemperatureMinTime,omitempty"`
	ApparentTemperatureMax     float64 `json:"apparentTemperatureMax,omitempty"`
	ApparentTemperatureMaxTime int     `json:"apparentTemperatureMaxTime,omitempty"`
	NearestStormBearing        float64 `json:"nearestStormBearing,omitempty,omitempty"`
	NearestStormDistance       float64 `json:"nearestStormDistance,omitempty,omitempty"`
	DewPoint                   float64 `json:"dewPoint,omitempty"`
	Humidity                   float64 `json:"humidity,omitempty"`
	WindSpeed                  float64 `json:"windSpeed,omitempty"`
	WindGust                   float64 `json:"windGust,omitempty"`
	WindGustTime               int     `json:"windGustTime,omitempty"`
	WindBearing                int     `json:"windBearing,omitempty"`
	Visibility                 float64 `json:"visibility,omitempty,omitempty"`
	CloudCover                 float64 `json:"cloudCover,omitempty"`
	Pressure                   float64 `json:"pressure,omitempty"`
	Ozone                      float64 `json:"ozone,omitempty"`
	UvIndex                    float64 `json:"uvIndex,omitempty"`
	UvIndexTime                int     `json:"uvIndexTime,omitempty"`
}

// DataBlock  contains the summary of the forcast and holds the datapoints
type DataBlock struct {
	Summary string      `json:"summary,omitempty"`
	Icon    string      `json:"icon,omitempty"`
	Data    []DataPoint `json:"data,omitempty"`
}

// Alert is the possible weather alerts in the area
type Alert struct {
	Title       string   `json:"title,omitempty"`
	Regions     []string `json:"regions,omitempty"`
	Severity    string   `json:"severity,omitempty"`
	Description string   `json:"description,omitempty"`
	Time        int64    `json:"time,omitempty"`
	Expires     float64  `json:"expires,omitempty"`
	URI         string   `json:"uri,omitempty"`
}

// Forecast is a full response from the api
type Forecast struct {
	Latitude     float64   `json:"latitude,omitempty"`
	Longitude    float64   `json:"longitude,omitempty"`
	Timezone     string    `json:"timezone,omitempty"`
	Offset       float64   `json:"offset,omitempty"`
	Currently    DataPoint `json:"currently,omitempty"`
	Minutely     DataBlock `json:"minutely,omitempty"`
	Hourly       DataBlock `json:"hourly,omitempty"`
	Daily        DataBlock `json:"daily,omitempty"`
	Alerts       []Alert   `json:"alerts,omitempty"`
	Flags        Flags     `json:"flags,omitempty"`
	APICalls     int       `json:"apiCalls,omitempty"`
	Code         int       `json:"code,omitempty"`
	ResponseTime int       `json:"responseTime,omitempty"`
}

// Units is an emum for the units type
type Units string

// Units types
const (
	CA   Units = "ca"
	SI   Units = "si"
	US   Units = "us"
	UK   Units = "uk"
	AUTO Units = "auto"
)

// Lang is an enum for the languages to recive back
type Lang string

// Lang Consts
const (
	Arabic             Lang = "ar"
	Azerbaijani        Lang = "az"
	Belarusian         Lang = "be"
	Bosnian            Lang = "bs"
	Catalan            Lang = "ca"
	Czech              Lang = "cs"
	German             Lang = "de"
	Greek              Lang = "el"
	English            Lang = "en"
	Spanish            Lang = "es"
	Estonian           Lang = "et"
	French             Lang = "fr"
	Croatian           Lang = "hr"
	Hungarian          Lang = "hu"
	Indonesian         Lang = "id"
	Italian            Lang = "it"
	Icelandic          Lang = "is"
	Cornish            Lang = "kw"
	Indonesia          Lang = "nb"
	Dutch              Lang = "nl"
	Polish             Lang = "pl"
	Portuguese         Lang = "pt"
	Russian            Lang = "ru"
	Slovak             Lang = "sk"
	Slovenian          Lang = "sl"
	Serbian            Lang = "sr"
	Swedish            Lang = "sv"
	Tetum              Lang = "te"
	Turkish            Lang = "tr"
	Ukrainian          Lang = "uk"
	IgpayAtinlay       Lang = "x-pig-latin"
	SimplifiedChinese  Lang = "zh"
	TraditionalChinese Lang = "zh-tw"
)

// GetForecast returns the forecast for a given lat long and time
func GetForecast(key string, lat string, long string, time string, units Units, lang Lang) (*Forecast, error) {
	res, err := getResponse(key, lat, long, time, units, lang)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	f, err := fromJSON(res.Body)
	if err != nil {
		return nil, err
	}

	calls, _ := strconv.Atoi(res.Header.Get("X-Forecast-API-Calls"))
	f.ResponseTime, _ = strconv.Atoi(res.Header.Get("X-Response-Time"))
	f.APICalls = calls

	return f, nil
}

func fromJSON(reader io.Reader) (*Forecast, error) {
	var f Forecast
	if err := json.NewDecoder(reader).Decode(&f); err != nil {
		return nil, err
	}

	return &f, nil
}

func getResponse(key string, lat string, long string, time string, units Units, lang Lang) (*http.Response, error) {
	coord := lat + "," + long

	var url string
	if time == "now" {
		url = BASEURL + "/" + key + "/" + coord + "?units=" + string(units) + "&lang=" + string(lang)
	} else {
		url = BASEURL + "/" + key + "/" + coord + "," + time + "?units=" + string(units) + "&lang=" + string(lang)
	}

	res, err := http.Get(url)
	if err != nil {
		return res, err
	}

	return res, nil
}
