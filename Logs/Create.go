package Logs

import (
	"fmt"
	"sync"

	"github.com/elucid503/Sprout-API-Go/Constants"
	"github.com/elucid503/Sprout-API-Go/Util"
)

type LogLevel int

const (

	// Log Levels

	LogLevelInfo    LogLevel = 0
	LogLevelWarning LogLevel = 1
	LogLevelError   LogLevel = 2
)

// Shared mutex for synchronizing log operations

var LogMutex sync.Mutex

func Log(ServiceUID string, Level LogLevel, Title string, Content string) {

	go func() {

		LogMutex.Lock()
		defer LogMutex.Unlock()

		Body := map[string]interface{}{
			"Level":   Level,
			"Title":   Title,
			"Content": Content,
		}

		// Send the log to the logging service

		Util.MakeHTTPRequest(fmt.Sprintf("%s/API/Services/%s/Logs/Create", Constants.LoggingRelayDomain, ServiceUID), "POST", map[string]string{

			"Content-Type": "application/json",

		}, Body)

	}()

}

func GetLogLevelString(Level LogLevel) string {

	switch Level {

		case LogLevelInfo:

			return "INFO"

		case LogLevelWarning:

			return "WARNING"

		case LogLevelError:

			return "ERROR"

		default:

			return "UNKNOWN"

	}

}