package geo

import (
	"encoding/json"
	"errors"
	"flyme-backend/app/config"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

var (
	ErrFailedToGetAccessToYOLP = errors.New("failed to get access to YOLP")
	ErrUnableToDecodeResponse  = errors.New("unable to decode YOLP response to JSON")
)

type Coordinate struct {
	Longitude float64
	Latitude  float64
}

func floatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func GetDistanceKm(coord []Coordinate) (float64, error) {

	if len(coord) <= 1 {
		return 0., nil
	}

	base := "https://map.yahooapis.jp/dist/V1/distance"
	appid := config.YOLP_APPID
	output := "json"

	coordsStr := ""
	for i, c := range coord {
		coordsStr += floatToStr(c.Longitude) + "," + floatToStr(c.Latitude)
		if i < len(coord)-1 {
			coordsStr += "%20"
		}
	}

	response, err :=
		http.Get(
			base + "?appid=" + appid + "&coordinates=" + coordsStr + "&output=" + output)

	if err != nil {
		return 0., ErrFailedToGetAccessToYOLP
	}

	var obj map[string]interface{}

	if err := json.NewDecoder(response.Body).Decode(&obj); err != nil {
		return 0., ErrUnableToDecodeResponse
	}

	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return 0., err
	}

	distance := gjson.Get(string(jsonStr), "Feature.0.Geometry.Distance").Float()

	return distance, nil
}
