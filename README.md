# search_stations_api

## Overview

API to list stations existing in any range from the selected point (represented by latitude and longitude).
Station data were obtained from https://ekidata.jp.

#### Request(Json)

```
{
  "limit_distance": number,
  "lon": number, //latitude
  "lat": number, //longitude
}
```

#### Response(Json)

```
{
  "selected_stations": 
    [
      {
        "station_name": string,
        "line_name": string,
        "distance: number, //Distance from the chosen point.
        },
       .
       .
       .
       {
        "station_nam": string,
        "line_name": string,
        "distance": number,
       }
     ]
}
```

## Build

```
docker-compose up -d --build
```

## Project setup

```
docker-compose exec app go run main.go
```

## API usage

- Get search reasults

```
curl http://localhost:8080/stations -X POST - "Content-Type: application/json" -d '{"limit_distance": 1.0, "lon": 139.698812, "lat": 35.68869}
```

## Technology used

- golang
- MySQL
- Docker

## Todo

- Introducing the rest
- Deploy
- Implement frontend
  - Get latitude and longitude using Google maps API
