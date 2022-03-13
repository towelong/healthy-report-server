package module

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Status string `json:"status"`
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		Precise    int    `json:"precise"`
		Confidence int    `json:"confidence"`
		Level      string `json:"level"`
	} `json:"result"`
}

func getLocationByAddress(address string) (*Location, error) {
	resp, err := http.Get("https://api.map.baidu.com/geocoder?address=" + address + "&output=json&key=f247cdb592eb43ebac6ccd27f796e2d2")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var l Location
	if resp.StatusCode == 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(b, &l)
		return &l, nil
	}
	return nil, errors.New("获取地址出错")
}
