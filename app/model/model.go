package model

type Station struct {
	Station_name string
	Line_name    string
	Line_cd      int
	Point        Point
}

type Point struct {
	Lon float64 //経度
	Lat float64 //緯度
}

type SelectedStations struct {
	Station_name string  `json:"station_name"`
	Line_name    string  `json:"line_name"`
	Distance     float64 `json:"distance"`
}

type Line struct {
	Line_cd   int
	Line_name string
}

type PointRequest struct {
	Limit_distance float64 `json:"limit_distance"`
	Lon            float64 `json:"lon"`
	Lat            float64 `json:"lat"`
}

type StationsResponse struct {
	SelectedStations []SelectedStations `json:"selected_stations"`
}
