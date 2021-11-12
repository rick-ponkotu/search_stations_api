package service

import (
	"app/model"
	"database/sql"
	"math"
	"sort"
)

type StationService struct {
	db *sql.DB
}

func NewStationService(db *sql.DB) *StationService {
	return &StationService{
		db: db,
	}
}

func (s *StationService) ReadLines() []model.Line {
	var lines []model.Line
	const sqlStr = `SELECT line_cd, line_name from app.table_line;`
	rows, err := s.db.Query(sqlStr)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var line model.Line
		err := rows.Scan(&line.Line_cd, &line.Line_name)
		if err != nil {
			panic(err.Error())
		}
		lines = append(lines, line)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	return lines
}

func (s *StationService) ReadStations(lines []model.Line) []model.Station {
	var stations []model.Station
	const sqlStr = `SELECT station_name,line_cd,lon,lat FROM app.table_station;`
	rows, err := s.db.Query(sqlStr)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var station model.Station
		err := rows.Scan(&station.Station_name, &station.Line_cd, &station.Point.Lon, &station.Point.Lat)
		if err != nil {
			panic(err.Error())
		}
		for _, line := range lines {
			if line.Line_cd == station.Line_cd {
				station.Line_name = line.Line_name
			}
		}
		stations = append(stations, station)
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	return stations
}

func Calc_distance(stations []model.Station, p model.Point, limit_distance float64) []model.SelectedStations {
	var selected_stations []model.SelectedStations
	for _, station := range stations {
		distance := calc(p, station.Point)
		if distance <= limit_distance {
			selected_station := model.SelectedStations{Station_name: station.Station_name, Line_name: station.Line_name, Distance: distance}
			selected_stations = append(selected_stations, selected_station)
		}
	}
	sort.Slice(selected_stations, func(i, j int) bool { return selected_stations[i].Distance < selected_stations[j].Distance })
	return selected_stations
}

func calc(s model.Point, t model.Point) float64 {
	const r = 6378.137
	const rad = math.Pi / 180
	delta_lon := t.Lon*rad - s.Lon*rad
	phi := math.Sin(s.Lat*rad)*math.Sin(t.Lat*rad) + math.Cos(s.Lat*rad)*math.Cos(t.Lat*rad)*math.Cos(delta_lon)
	d := r * math.Acos(phi)
	return d
}
