package Logging

import (
	"fmt"

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

func Log(ServiceUID string, Level LogLevel, Title string, Content string) error {

	Body := map[string]interface{}{

		"Level":   Level,
		"Title":   Title,
		"Content": Content,
	}

	// Send the log to the logging service

	_, Error := Util.MakeHTTPRequest(fmt.Sprintf("%s/API/Services/%s/Logs/Create", Constants.LoggingRelayDomain, ServiceUID), "POST", map[string]string{

		"Content-Type": "application/json",

	}, Body)

	return Error

}