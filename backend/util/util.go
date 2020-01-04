package storkutil

import (
	"os"
	"time"
	"fmt"
	"path"
	"runtime"
	log "github.com/sirupsen/logrus"
)


func UTCNow() time.Time {
	return time.Now().UTC()
}

// Returns URL of the localhost with port.
func LocalHostWithPort(port int64) string {
	return fmt.Sprintf("http://localhost:%d/", port)
}

func SetupLogging() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
		// TODO: do more research and enable if it brings value
		//PadLevelText: true,
		// FieldMap: log.FieldMap{
		// 	FieldKeyTime:  "@timestamp",
		// 	FieldKeyLevel: "@level",
		// 	FieldKeyMsg:   "@message",
		// },
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// Grab filename and line of current frame and add it to log entry
			_, filename := path.Split(f.File)
			return "", fmt.Sprintf("%20v:%-5d", filename, f.Line)
		},
	})
}