package crmlogutil

import (
	"encoding/json"
	"fmt"
	"go/build"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var sourceHostName string
var messageBean = LogMessageBean{}
var goPath = build.Default.GOPATH

func init() {
	initialHostName()
}

type SetLogger struct {
	IsJSON    bool   `json:"isJson"`
	WriteFile bool   `json:"writeFile"`
	Path      string `json:"path"`
	FileName  string `json:"fileName"`
}

type PatternLogger struct {
	ApplicationName string
	ProductName     LogProductName
	SourceSystem    LogSystem
	TargetSystem    LogSystem
	SetLogger       SetLogger
}

func initialHostName() {
	if sourceHostName == "" {
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		sourceHostName = hostname
	}
}

func getCallerInfoBean(stacktrace StackTrace) (callerInfoBean CallerInfoBean) {
	runes := []rune(stacktrace.Function)
	indexFirst := strings.Index(stacktrace.Function, ".")
	lastChar := string(stacktrace.Function[len(stacktrace.Function)-1:])
	indexLast := strings.LastIndex(stacktrace.Function, lastChar)
	callerInfoBean.ClassName = string(runes[0:indexFirst])
	callerInfoBean.FileName = fmt.Sprintf("\n\t %s:%d", stacktrace.File, stacktrace.Line)
	callerInfoBean.MethodName = string(runes[indexFirst+1 : indexLast+1])
	return callerInfoBean
}

func getElapsedTime(start time.Time) int64 {
	return time.Now().Sub(start).Nanoseconds() / int64(time.Millisecond)
}

func (p *PatternLogger) logMonitoring(messageType string, correlationID string, monitorType LogMonitorType, targetURL string,
	action string, elapsedTime int64, responseCode string) string {

	messageBean.Timestamp = time.Now().Format(time.RFC3339Nano)
	messageBean.ApplicationName = p.ApplicationName
	messageBean.LogType = Monitor
	messageBean.CorrelationID = correlationID
	messageBean.MessageType = messageType
	messageBean.ProductName = p.ProductName
	messageBean.MonitorType = monitorType
	messageBean.SourceSystem = p.SourceSystem
	messageBean.SourceHostName = sourceHostName
	messageBean.TargetSystem = p.TargetSystem
	messageBean.TargetURL = targetURL
	messageBean.Action = action
	messageBean.ElapsedTime = elapsedTime
	messageBean.ResponseCode = responseCode
	if !p.SetLogger.IsJSON {
		if !p.SetLogger.WriteFile {
			fmt.Fprintln(os.Stdout, logStringPattern(&messageBean))
		} else {
			p.createLogFile(logStringPattern(&messageBean))
		}
	} else {
		if !p.SetLogger.WriteFile {
			enc := json.NewEncoder(os.Stdout)
			enc.SetEscapeHTML(false)
			enc.SetIndent("", "    ")
			enc.Encode(messageBean)
		} else {
			marshalStruct, _ := json.Marshal(&messageBean)
			p.createLogFile(string(marshalStruct))
		}
	}

	return messageBean.Timestamp
}

func (p *PatternLogger) LogRequest(correlationID string, monitorType LogMonitorType, targetURL string, action string) string {
	return p.logMonitoring("Request", correlationID, monitorType, targetURL, action, getElapsedTime(time.Time{}), "")
}

func (p *PatternLogger) LogRequestApplication(correlationID string) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Request", correlationID, Application, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(time.Time{}), "")
}

func (p *PatternLogger) LogRequestXMLProvider(correlationID string) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Request", correlationID, XMLProvider, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(time.Time{}), "")
}

func (p *PatternLogger) LogRequestXMLClient(correlationID string, targetURL string, action string) string {
	return p.logMonitoring("Request", correlationID, XMLClient, targetURL, action, getElapsedTime(time.Time{}), "")
}

func (p *PatternLogger) LogRequestWSProvider(correlationID string) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Request", correlationID, WebServiceProvider, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(time.Time{}), "")
}

func (p *PatternLogger) LogRequestWSClient(correlationID string, targetURL string, action string) string {
	return p.logMonitoring("Request", correlationID, WebServiceClient, targetURL, action, getElapsedTime(time.Time{}), "")
}

func (p *PatternLogger) LogRequestRESTProvider(correlationID string) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Request", correlationID, RESTServiceProvider, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(time.Time{}), "")
}

func (p *PatternLogger) LogRequestRESTClient(correlationID string, targetURL string, action string) string {
	return p.logMonitoring("Request", correlationID, RESTServiceClient, targetURL, action, getElapsedTime(time.Time{}), "")
}

func (p *PatternLogger) LogRequestDBClient(correlationID string) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Request", correlationID, DatabaseClient, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(time.Time{}), "")
}

func (p *PatternLogger) LogResponse(correlationID string, monitorType LogMonitorType, targetURL string, action string,
	responseCode string, requestDateTime time.Time) string {
	return p.logMonitoring("Response", correlationID, monitorType, targetURL, action, getElapsedTime(requestDateTime), responseCode)
}

func (p *PatternLogger) LogResponseApplication(correlationID string, responseCode string, requestDateTime time.Time) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Response", correlationID, Application, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(requestDateTime), responseCode)
}

func (p *PatternLogger) LogResponseXMLProvider(correlationID string, responseCode string, requestDateTime time.Time) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Response", correlationID, XMLProvider, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(requestDateTime), responseCode)
}

func (p *PatternLogger) LogResponseXMLClient(correlationID string, targetURL string, action string, responseCode string,
	requestDateTime time.Time) string {
	return p.logMonitoring("Response", correlationID, XMLClient, targetURL, action, getElapsedTime(requestDateTime), responseCode)
}

func (p *PatternLogger) LogResponseWSProvider(correlationID string, responseCode string, requestDateTime time.Time) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Response", correlationID, WebServiceProvider, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(requestDateTime), responseCode)
}

func (p *PatternLogger) LogResponseWSClient(correlationID string, targetURL string, action string, responseCode string,
	requestDateTime time.Time) string {
	return p.logMonitoring("Response", correlationID, WebServiceClient, targetURL, action, getElapsedTime(requestDateTime), responseCode)
}

func (p *PatternLogger) LogResponseRESTProvider(correlationID string, responseCode string, requestDateTime time.Time) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Response", correlationID, RESTServiceProvider, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(requestDateTime), responseCode)
}

func (p *PatternLogger) LogResponseRESTClient(correlationID string, targetURL string, action string, responseCode string, requestDateTime time.Time) string {
	return p.logMonitoring("Response", correlationID, RESTServiceClient, targetURL, action, getElapsedTime(requestDateTime), responseCode)
}

func (p *PatternLogger) LogResponseDBClient(correlationID string, responseCode string, requestDateTime time.Time) string {
	callerInfoBean := getCallerInfoBean(GetStackTrace())

	return p.logMonitoring("Response", correlationID, DatabaseClient, callerInfoBean.ClassName,
		callerInfoBean.MethodName, getElapsedTime(requestDateTime), responseCode)
}

func (p *PatternLogger) logApp(correlationID string, level LogLevel, message string, args ...interface{}) {
	messageBean.Timestamp = time.Now().Format(time.RFC3339Nano)
	messageBean.ApplicationName = p.ApplicationName
	messageBean.LogType = AppLog
	messageBean.CorrelationID = correlationID
	messageBean.Level = level
	messageBean.Message = message
	messageBean.StackTrace = args[0]
	messageBean.SourceHostName = sourceHostName
	messageBean.ProductName = p.ProductName
	messageBean.SourceSystem = p.SourceSystem
	messageBean.TargetSystem = p.TargetSystem

	if !p.SetLogger.IsJSON {
		if !p.SetLogger.WriteFile {
			fmt.Fprintln(os.Stdout, logStringPattern(&messageBean))
		} else {
			p.createLogFile(logStringPattern(&messageBean))
		}
	} else {
		if !p.SetLogger.WriteFile {
			enc := json.NewEncoder(os.Stdout)
			enc.SetEscapeHTML(false)
			enc.SetIndent("", "    ")
			enc.Encode(messageBean)
		} else {
			marshalStruct, _ := json.Marshal(&messageBean)
			p.createLogFile(string(marshalStruct))
		}
	}
}

func (p *PatternLogger) WriteRequestMsg(correlationID string, url string, message interface{}) {
	if message == nil {
		p.logApp(correlationID, INFO, "Request URL: "+url+", Request Message: null", nil)
	} else {
		msg, err := json.Marshal(message)
		if err != nil {
			p.Error(correlationID, err.Error(), Wrap(err))
		}
		p.logApp(correlationID, INFO, "Request URL: "+url+", Request Message: "+string(msg), nil)
	}
}

func (p *PatternLogger) WriteResponseMsg(correlationID string, message interface{}) {
	if message == nil {
		p.logApp(correlationID, INFO, "Request Message: null", nil)
	} else {
		msg, err := json.Marshal(message)
		if err != nil {
			p.Error(correlationID, err.Error(), Wrap(err))
		}
		p.logApp(correlationID, INFO, "Request Message: "+string(msg), nil)
	}
}

func (p *PatternLogger) Info(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, INFO, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, INFO, message, nil)
	}
}

func (p *PatternLogger) Fatal(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, FATAL, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, FATAL, message, nil)
	}
}

func (p *PatternLogger) Error(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, ERROR, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, ERROR, message, nil)
	}
}

func (p *PatternLogger) Warn(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, WARN, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, WARN, message, nil)
	}
}

func (p *PatternLogger) Debug(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, DEBUG, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, DEBUG, message, nil)
	}
}

func (p *PatternLogger) Trace(correlationID string, message string, args ...interface{}) {
	if len(args) > 0 {
		msg, stackTrace := checkArguments(args)
		p.logApp(correlationID, TRACE, message+msg, stackTrace)
	} else {
		p.logApp(correlationID, TRACE, message, nil)
	}
}

func checkArguments(args []interface{}) (message string, staceTrace string) {
	for i := 0; i < len(args); i++ {
		switch v := args[i].(type) {
		default:
			str, _ := json.Marshal(v)
			message = message + " " + string(str)
		case float32, float64, complex64, complex128:
			message = message + fmt.Sprintf("%s %g", " ", v)
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			message = message + fmt.Sprintf("%s %d", " ", v)
		case bool:
			message = message + fmt.Sprintf("%s %t", " ", v)
		case string:
			message = message + " " + v
		case *strconv.NumError, Error:
			staceTrace = staceTrace + fmt.Sprintf("%s %v", " ", v)
		}
	}
	return
}

func logStringPattern(bean *LogMessageBean) (log string) {
	log = fmt.Sprintf(`timestamp=%s|logType=%s|correlationID=%s|level=%v|message=%s|messageType=%s|stackTrace=%v|productName=%s|monitorType=%v|sourceSystem=%s|sourceHostName=%s|targetSystem=%s|targetURL=%s|action=%s|elapsedTime=%d|responseCode=%s`,
		bean.Timestamp, bean.LogType, bean.CorrelationID, bean.Level, bean.Message, bean.MessageType, bean.StackTrace, bean.ProductName,
		bean.MonitorType, bean.SourceSystem, bean.SourceHostName, bean.TargetSystem, bean.TargetURL, bean.Action, bean.ElapsedTime, bean.ResponseCode)
	return
}

func (p *PatternLogger) createLogFile(logging string) {
	file, _ := os.OpenFile(goPath+p.SetLogger.Path+"/"+p.SetLogger.FileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	log.SetOutput(file)
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.SetFlags(0)
	log.Println(logging)
}
