<img src="gostsite/resources/img/icon.png" width="353">  
<a href="http://beta.drone.io/drone/drone"><img src="http://beta.drone.io/api/badges/drone/drone/status.svg" /></a>
[![Coverage Status](https://coveralls.io/repos/github/Geodan/gost/badge.svg?branch=master)](https://coveralls.io/github/Geodan/gost?branch=master)

GOST (Go-SensorThings) is a sensor server written in Go. It implements the [OGC SensorThings API] (http://ogc-iot.github.io/ogc-iot-api/api.html) standard.

## Disclaimer

GOST is alpha software and is not considered appropriate for customer use. Feel free to help development :-)

## License

GOST licensed under [MIT](https://opensource.org/licenses/MIT).

## Getting started

1] Install
Install GoLang https://golang.org/

Install Postgresql http://www.postgresql.org/

2] Clone code
git clone https://github.com/Geodan/gost.git

3] Get dependencies

go get gopkg.in/yaml.v2

go get github.com/lib/pq 

go get github.com/gorilla/mux

4] Edit config.yaml
Change connection to database

5] Start
go run main.go

## Dependencies

[yaml v2](https://github.com/go-yaml/yaml)
[pq](https://github.com/lib/pq)  
[mux](https://github.com/gorilla/mux)  
[SurgeMQ](github.com/surgemq/surgemq)  

## Roadmap

- Complete implementation of the OGC SensorThings spec
- Tests!
- MQTT
- Frontend
- Benchmarks
- Different storage providers such as MongoDB (Now using PostgreSQL)

## TODO

[See wiki](https://github.com/Geodan/gost/wiki/TODO)
