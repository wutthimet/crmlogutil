package crmlog

type monitorTypes string

func (monitorType monitorTypes) isMonitorType() monitorTypes {
	return monitorType
}

type LogMonitorType interface {
	isMonitorType() monitorTypes
}

const (
	Application   = monitorTypes("Application")
	XMLProvider  = monitorTypes("XMLProvider")
	XMLClient = monitorTypes("XMLClient")
	WebServiceProvider = monitorTypes("WebServiceProvider")
	WebServiceClient   = monitorTypes("WebServiceClient")
	RESTServiceProvider = monitorTypes("RESTServiceProvider")
	RESTServiceClient   = monitorTypes("RESTServiceClient")
	DatabaseClient  = monitorTypes("DatabaseClient")
)