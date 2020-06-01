package scheduler

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sahilgarg96/DBTNT/handler"
	"github.com/sahilgarg96/DBTNT/logging"
	"github.com/sahilgarg96/DBTNT/redis"
	"strconv"
	"strings"
	"time"
)

var Logger = logging.NewLogger()

func Init() {

	schedule := gocron.NewScheduler()
	schedule.Every(1).Seconds().Do(Task)
	schedule.Start()
}

func Task() {
	//get all keys
	keys := redis.Keys("pdftsuser_*")

	for i := range keys.Val() {

		key := keys.Val()[i]

		value, err := redis.GetValue(key)

		if err != nil {
			Logger.Errorf("Error getting key " + err.Error())
		}

		t1, errr := strconv.ParseInt(value, 10, 64)
		if errr != nil {
			panic(errr)
		}
		//parse current time
		t2 := time.Now().Unix()
		diff := t2 - t1

		if diff >= 300 { // 5mins

			userId := strings.Split(key, "_")[1]
			err = handler.SendEmail("support@dountnut.com", "username@gmail.com", userId+"_"+value+".pdf")
			if err != nil {
				Logger.Errorf("error sending mail " + err.Error())
			}
			redis.Del(key)
		}

	}

}
