package handler

import (
	"app/model"
	"app/service"
	"encoding/json"
	"net/http"
)

type StationHandler struct {
	svc *service.StationService
}

func NewStationHandler(svc *service.StationService) *StationHandler {
	return &StationHandler{
		svc: svc,
	}
}

func (h *StationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	var p_req model.PointRequest
	if err := d.Decode(&p_req); err != nil {
		panic(err.Error())
	}
	lines := h.svc.ReadLines()
	stations := h.svc.ReadStations(lines)
	req := model.PointRequest{Limit_distance: p_req.Limit_distance, Lon: p_req.Lon, Lat: p_req.Lat}
	p := model.Point{Lon: req.Lon, Lat: req.Lat}

	Selected_Stations := service.Calc_distance(stations, p, req.Limit_distance)
	var response model.StationsResponse
	for _, selected_station := range Selected_Stations {
		response.SelectedStations = append(response.SelectedStations, selected_station)
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		panic(err.Error())
	}

}
