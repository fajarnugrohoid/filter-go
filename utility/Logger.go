package utility

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var lock = &sync.Mutex{}
var singleInstance *logrus.Logger

var Arg1 *string

func InstanceLogger(arg string) *logrus.Logger {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = logrus.New()

			// get the location
			location, _ := time.LoadLocation("Asia/Jakarta")
			// this should give you time in location
			t := time.Now().In(location)
			fmt.Println(t)
			t.Year()
			t.Month()
			t.Day()
			t.Hour()
			t.Minute()
			t.Second()
			//formatted := fmt.Sprintf("%d%02d%02dT%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

			//filename := "logs/log_" + arg + "_" + formatted + ".log"
			filename := "logs/log_" + arg + "_" + ".log"
			file, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)

			val := os.Getenv("STACK")
			fmt.Println(val)
			if os.Getenv("LOGGING") == "file" {
				singleInstance.SetOutput(file)
				singleInstance.SetLevel(logrus.DebugLevel)
			}

		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func GetLogArg() string {
	if Arg1 != nil {
		return *Arg1
	}
	return ""

}

func SetLogArg(val string) {
	Arg1 = &val
}
