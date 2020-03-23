package influxdb

import (
	"errors"
	"melody/config"
	"time"
)

const (
	Namespace = "melody_influxdb"
	defaultDB = "melody_data"
	defaultAddress = "*:8086"
)

var configErr = errors.New("load influx config error")

// influxdbConfig 描述metrics输出的influxDB的信息
type influxdbConfig struct {
	address    string
	username   string
	password   string
	ttl        time.Duration
	db         string
	bufferSize int
	timeout    time.Duration
}

func getConfig(config config.ExtraConfig) interface{} {
	v, ok := config[Namespace]
	if !ok {
		return nil
	}

	mapStruct, ok := v.(map[string]interface{})
	if !ok {
		return nil
	}

	influx := influxdbConfig{}

	if value, ok := mapStruct["address"]; ok {
		influx.address = value.(string)
	} else {
		influx.address = defaultAddress
	}

	if value, ok := mapStruct["username"]; ok {
		influx.username = value.(string)
	}

	if value, ok := mapStruct["password"]; ok {
		influx.password = value.(string)
	}

	if value, ok := mapStruct["db"]; ok {
		influx.db = value.(string)
	} else {
		influx.db = defaultDB
	}

	if size, ok := mapStruct["buffer_size"]; ok {
		if s, ok := size.(int); ok {
			influx.bufferSize = s
		}
	}

	if ttl, ok := mapStruct["ttl"]; ok {
		t, ok := ttl.(string)
		if !ok {
			return nil
		}
		var err error
		influx.ttl, err = time.ParseDuration(t)
		if err != nil {
			return nil
		}
	}

	if timeout, ok := mapStruct["time_out"]; ok {
		t, ok := timeout.(string)
		if !ok {
			return nil
		}
		var err error
		influx.timeout, err = time.ParseDuration(t)
		if err != nil {
			return nil
		}
	}

	return influx
}