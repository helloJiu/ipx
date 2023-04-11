package ipx

import (
	"net"

	"github.com/oschwald/geoip2-golang"
)

type Address struct {
	City      string
	Province  string
	Country   string
	Continent string
	TimeZone  string
	Latitude  float64
	Longitude float64
}

var db *geoip2.Reader

func Query(ipAddr string) (*Address, error) {
	ip := net.ParseIP(ipAddr)
	record, err := db.City(ip)
	if err != nil {
		return nil, err
	}

	var addr Address
	// 城市名称
	if len(record.City.Names) > 0 {
		addr.City = record.City.Names["zh-CN"]
	}

	// 省份
	if len(record.Subdivisions) > 0 {
		if len(record.Subdivisions[0].Names) > 0 {
			addr.Province = record.Subdivisions[0].Names["zh-CN"]
		}
	}

	// 国家名
	if len(record.Country.Names) > 0 {
		addr.Country = record.Country.Names["zh-CN"]
	}

	// // 洲名
	// if len(record.Continent.Names) > 0 {
	// 	addr.Cib = record.Continent.Names["zh-CN"]
	// }

	// 时区
	addr.TimeZone = record.Location.TimeZone

	// 纬度
	addr.Latitude = record.Location.Latitude

	// 经度
	addr.Longitude = record.Location.Longitude
	return &addr, nil
}

func InitIpx(dataPath string) error {
	var err error
	db, err = geoip2.Open(dataPath)
	return err
}

func CloseDB() {
	db.Close()
}
