package logging

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

// n = name of log category
func InitLogger(n string) *log.Logger {
	var logBuffer bytes.Buffer
	var logFile *os.File
	var logger *log.Logger
	var err error

	logFile, err = os.OpenFile("./logs/"+n+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("! Cannot open "+n+" log for writing:", err)
		return nil
	}

	logger = log.New(&logBuffer, n+": ", log.Lshortfile)
	logger.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " - " + n + ": ")
	logger.SetOutput(logFile)

	return logger
}

/*func CloseLogger(l *log.Logger) {

}*/
