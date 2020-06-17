package crmlogutil

type logTypes string

func (logType logTypes) isType() logTypes {
	return logType
}

type LogType interface {
	isType() logTypes
}

const (
	Monitor   = logTypes("Monitor")
	AppLog    = logTypes("AppLog")
	ServerLog = logTypes("ServerLog")
)
