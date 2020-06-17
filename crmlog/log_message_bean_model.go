package crmlog

type LogMessageBean struct {
	Timestamp       string         `json:"@timestamp""`
	ApplicationName string         `json:"@suffix"`
	LogType         LogType        `json:"logType"`
	CorrelationID   string         `json:"correlationID"`
	Level           LogLevel       `json:"level"`
	Message         string         `json:"message"`
	MessageType     string         `json:"messageType"`
	StackTrace      interface{}    `json:"stackTrace"`
	ProductName     LogProductName `json:"productName"`
	MonitorType     LogMonitorType `json:"monitorType"`
	SourceSystem    LogSystem      `json:"sourceSystem"`
	SourceHostName  string         `json:"sourceHostName"`
	TargetSystem    LogSystem      `json:"targetSystem"`
	TargetURL       string         `json:"targetURL"`
	Action          string         `json:"action"`
	ElapsedTime     int64          `json:"elapsedTime"`
	ResponseCode    string         `json:"responseCode"`
}
