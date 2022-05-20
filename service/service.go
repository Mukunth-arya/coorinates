package service

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	data "github.com/Mukunth-arya/coorinates/proto/coordinates"
)

const (
	Map_Coordinates = "https://maps.googleapis.com/maps/api/geocode/json?key=MYKEY&address="
)

type Dataresult struct {
	Result Results
}
type Results []Geometry
type Geometry struct {
	Geometry Location
}
type Location struct {
	Location Coordinates
}
type Coordinates struct {
	Latitute  string
	Longitude string
}

func Coordinatedata(ctx context.Context, rr *data.CordinateRequest) (*data.CoordinateResponse, error) {

	resp, err := http.Get(Map_Coordinates + rr.City)

	if err != nil {
		return &data.CoordinateResponse{}, errors.New("Please enter the proper city name")
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &data.CoordinateResponse{}, errors.New("Unable to read the Value")
	}
	var datas Dataresult
	json.Unmarshal(bytes, &datas)

	return &data.CoordinateResponse{CoordinateData: datas.Result[0].Geometry.Location.Latitute}, errors.New("Sorry unable to process th request!!!!")

}
