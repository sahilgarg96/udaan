package models

import (
	"time"
)

type User struct {
	UserId int32
	Name   string
	Phone  string
	Email  string
}

type Device struct {
	DeviceId   int32
	DeviceName string
	DeviceType int32
}

type Metrics struct {
	MetricId          int32
	MetricName        string
	MetricDefaultUnit string
}

type DeviceAvailableMetrics struct {
	DeviceAvailableMetricId int32
	DeviceId                int32
	MetricId                int32
	DeletedFlag             int32
}

type UserDeviceMapping struct {
	UserDeviceMappingId int32
	UserId              int32
	DeviceId            int32
}

type UserDeviceMetricData struct {
	MetricDataId        int32
	UserDeviceMappingId int32
	UserId              int32
	DeviceId            int32
	MetricValue         int32
	CreatedAt           time.Time
	DeletedFlag         int32
}
