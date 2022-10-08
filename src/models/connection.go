package models

type Connection struct {
	IP        string `json:"ip"`
	ISP       string `json:"isp"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
