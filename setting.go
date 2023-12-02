package main

import (
	"net/url"
	"strconv"
	"syscall/js"
)

type Setting struct {
	goidsNum int
	maxSpeed float64
	maxForce float64
	sight    float64
}

func NewSetting() Setting {
	setting := Setting{goidsNum: 30, maxSpeed: 4, maxForce: 2, sight: 100}
	if goidsNum, err := getQueryValueInt("num"); err == nil {
		setting.goidsNum = goidsNum
	}
	if maxSpeed, err := getQueryValueFloat("speed"); err == nil {
		setting.maxSpeed = maxSpeed
	}
	if maxForce, err := getQueryValueFloat("force"); err == nil {
		setting.maxForce = maxForce
	}
	if sight, err := getQueryValueFloat("sight"); err == nil {
		setting.sight = sight
	}
	return setting
}

func getURL() string {
	return js.Global().Get("location").Get("href").String()
}

func getQuery() url.Values {
	u, _ := url.Parse(getURL())
	return u.Query()
}

func getQueryValue(key string) string {
	return getQuery().Get(key)
}

func getQueryValueInt(key string) (int, error) {
	return strconv.Atoi(getQueryValue(key))
}

func getQueryValueFloat(key string) (float64, error) {
	return strconv.ParseFloat(getQueryValue(key), 64)
}
