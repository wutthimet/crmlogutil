package crmlogutil

type logLevels string

func (logLevel logLevels) isLevel() logLevels {
	return logLevel
}

type LogLevel interface {
	isLevel() logLevels
}

const (
	FATAL = logLevels("FATAL")
	ERROR = logLevels("ERROR")
	WARN  = logLevels("WARN")
	INFO  = logLevels("INFO")
	DEBUG = logLevels("DEBUG")
	TRACE = logLevels("TRACE")
	ALL   = logLevels("ALL")
)
