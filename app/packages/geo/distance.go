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

func floatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func GetDistanceKm(alng, alat, blng, blat float64) (float64, error) {
	base := "https://map.yahooapis.jp/dist/V1/distance"
	appid := config.YOLP_APPID
	output := "json"

	response, err :=
		http.Get(
			base + "?appid=" + appid + "&coordinates=" +
				floatToStr(alng) + "," +
				floatToStr(alat) + "%20" +
				floatToStr(blng) + "," +
				floatToStr(blat) + "&output=" + output)

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
