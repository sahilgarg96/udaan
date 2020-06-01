package main

import (
	"bufio"
	"fmt"
	"github.com/sahilgarg96/udaan/app"
	"github.com/sahilgarg96/udaan/handler"
	"github.com/sahilgarg96/udaan/logging"
	"io"
	"log"
	"net/http"
)

var Logger = logging.NewLogger()

const (
	adduser      = "ADDU"
	adddevice    = "ADDD"
	pushmetric   = "PUSHM"
	getmetric    = "GETM"
	getavgmetric = "GETAM"
)

func fetchCommands(a *app.AppData) {

	input := bufio.Newreader(os.Stdin)

	for {

		ip, err := input.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			break
		}

		it := strings.Replace(ip, "\n", "", -1)
		params := strings.Split(it, " ")
		coomand := params[0]

		err = executeCommand(coomand, params, a)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func executeCommand(command string, params []string, a *app.AppData) {

	switch command {
	case "ADDU":
		
		a.AddUser(params[1],params[2],params[3])
	break
	case "ADDD"
		a.AddDeviceForAUser(params[1],params[2])
	break
	case "PUSHM"
		a.AddUserMetric(params[1],params[2],params[3], params[4])

	break

	case "GETM"
		a.GetAllDataForMetricAndUser(params[1],params[2])

	break

	case "GETAM"

		a.GetSpecifiedDataForMetricAndUser(params[1],params[2])
	break
	}

	fmt.Println(a)
}

func main() {

	Logger.Infof("starting service ")

	newApp := &app.AppData{}

	newApp.intializeApp()

	fetchCommands(newApp)

	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
	if err != nil {
		log.Fatal("some error occured")
	}

}
