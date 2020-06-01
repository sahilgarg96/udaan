package structs

import "github.com/sahilgarg96/udaan/models/models"

import "strconv"

type AppData struct {
	Users                  *[]models.User
	Devices                *[]models.Device
	Metrics                *[]models.Metrics
	DeviceAvailableMetrics *[]models.DeviceAvailableMetrics
	UserDeviceMappings     *[]models.UserDeviceMapping
	UserDeviceMetricData   *[]models.UserDeviceMetricData
}

func (a *AppData) AddDefaultUsers() {

	a.AddUser("Sahil", "88839292992", "sahdhh@hddj.com")
	a.AddUser("Nikhil", "88859292992", "aahdhh@hddj.com")
	a.AddUser("Raman", "98839292992", "zahdhh@hddj.com")
	a.AddUser("Kshitij", "78839292992", "xahdhh@hddj.com")
}

func (a *AppData) AddDefaultDevices() {

	a.AddDevice("1", "Fitbit")
	a.AddDevice("2", "MiFit")
	a.AddDevice("3", "Strava")
	a.AddDevice("4", "AppleWatch")
	a.AddDevice("5", "SamsungWatch")
}

func (a *AppData) AddUser(name string, phone string, email string) {

	var uid int32
	if len(a.Users) > 0 {
		uid = a.Users[len(a.Users)-1].UserId + 1
	} else {
		uid = 1
	}

	user := models.User{
		UserId: uid,
		Name:   name,
		Phone:  phone,
		Email:  email,
	}

	a.Users = append(a.Users, user)

	fmt.Println("Added user succesfully")

	fmt.Println(a.Users)
}

func (a *AppData) AddDevice(deviceId string, name string) {

	did, _ := strconv.Atoi(deviceId)
	device := models.Device{
		DeviceId:   did,
		DeviceName: name,
		DeviceType: 1,
	}

	a.Devices = append(a.Devices, device)
}

func (a *AppData) AddDeviceForAUser(deviceId string, userid string) {

	var udm_id int32

	if len(a.UserDeviceMappings) > 0 {
		udm_id = a.UserDeviceMappings[len(a.UserDeviceMappings)-1].UserDeviceMappingId + 1
	} else {
		udm_id = 1
	}

	did, _ := strconv.Atoi(deviceId)
	uid, _ := strconv.Atoi(userId)
	udm := models.UserDeviceMapping{
		UserId:              uid,
		DeviceId:            did,
		UserDeviceMappingId: udm_id,
	}

	a.UserDeviceMapping = append(a.UserDeviceMapping, device)

	fmt.Println("Added device succesfully")

	fmt.Println(a.UserDeviceMapping)
}

func (a *AppData) AddMetrics() {

	metric = models.Metric{
		MetricId:          1,
		MetricName:        "Height",
		MetricDefaultUnit: "inches",
	}
	a.Metrics = append(a.Metrics, metric)

	metric = models.Metric{
		MetricId:          2,
		MetricName:        "Weight",
		MetricDefaultUnit: "kg",
	}
	a.Metrics = append(a.Metrics, metric)

	metric = models.Metric{
		MetricId:          3,
		MetricName:        "HeartRate",
		MetricDefaultUnit: "bpm",
	}
	a.Metrics = append(a.Metrics, metric)

	metric = models.Metric{
		MetricId:          4,
		MetricName:        "CalorieBurn",
		MetricDefaultUnit: "cal",
	}
	a.Metrics = append(a.Metrics, metric)
}

func (a *AppData) IntializeApp() {

	a.AddMetrics()
	a.AddDefaultDevices()

	itr := 1

	for _, device := range a.Devices {

		for _, metric := range a.Metrics {

			dev_met := models.DeviceAvailableMetrics{
				DeviceAvailableMetricId: itr,
				DeviceId:                device.DeviceId,
				MetricId:                device.MetricId,
				DeletedFlag:             0,
			}

			itr++
		}
	}

}

func (a *AppData) GetUserDeviceMappingId(userId, deviceId int32) int32 {

	var id int32
	for _, row := range a.UserDeviceMappings {

		if row.UserId == userId && row.DeviceId == deviceId {
			id = row.UserDeviceMappingId
			break
		}
	}

	return id
}

func (a *AppData) AddUserMetric(user_id, device_id, metric_id, value string) {

	userId, _ := strconv.Atoi(user_id)
	deviceId, _ := strconv.Atoi(device_id)
	metric_id, _ := strconv.Atoi(metric_id)
	value, _ := strconv.Atoi(value)

	var id int32

	if len(a.UserDeviceMetricData) > 0 {
		id = a.UserDeviceMetricData[len(a.UserDeviceMetricData)-1].MetricDataId + 1
	} else {
		id = 1
	}

	c_time := time.Now()

	udm_id := a.GetUserDeviceMappingId(userId, deviceId)
	data := models.UserDeviceMetricData{
		MetricDataId:        id,
		UserId:              userId,
		DeviceId:            deviceId,
		UserDeviceMappingId: udm_id,
		MetricValue:         value,
		CreatedAt:           c_time,
		DeletedFlag:         0,
	}

	fmt.Println("Added metric succesfully")
}

func (a *AppData) GetAllDataForMetricAndUser(user_id, metric_id string) {

	userId, _ := strconv.Atoi(user_id)
	metric_id, _ := strconv.Atoi(metric_id)

	for _, row := range a.UserDeviceMetricData {

		if row.UserId == userId && row.MetricId == metric_id {

			fmt.Println(row)
		}
	}
}

func (a *AppData) GetSpecifiedDataForMetricAndUser(user_id, metric_id string) {

	userId, _ := strconv.Atoi(user_id)
	metric_id, _ := strconv.Atoi(metric_id)

	var min, max, average int32

	var dataAvailable bool

	min := 1000000000
	max := 0
	average := 0

	cumm := 0
	ct := 0
	for _, row := range a.UserDeviceMetricData {

		if row.UserId == userId && row.MetricId == metric_id {

			dataAvailable = true
			cumm += row.MetricValue
			ct++
			if row.MetricValue > max {
				max = row.MetricValue
			}

			if row.MetricValue < min {
				min = row.MetricValue
			}
		}
	}
	if ct > 0 {
		average = cumm / ct
	}

	if !dataAvailable {
		fmt.Println("No data available")
	} else {
		fmt.Println("Max", max, "         ")
		fmt.Println("Min", min, "        ")
		fmt.Println("Average", average, "       ")
	}
}
