package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Initialise Logger with default values.
// Logger will write logs in "fizzbuzz.log" file.
// You can set the log level using LOG_LEVEL={TRACE,DEBUG,INFO,WARN} environment variable.
// default value : INFO
func Init() {
	val, present := os.LookupEnv("LOG_LEVEL")
	if !present {
		log.SetLevel(log.InfoLevel)
	} else {
		switch val {
		case "TRACE":
			log.SetLevel(log.TraceLevel)
		case "DEBUG":
			log.SetLevel(log.DebugLevel)
		case "WARN":
			log.SetLevel(log.WarnLevel)
		default:
			log.SetLevel(log.InfoLevel)
		}
	}

	logfile, err := os.OpenFile("fizzbuzz.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(logfile)
	} else {
		log.Warnln("Can not log in file ! Using std output.")
	}

}
