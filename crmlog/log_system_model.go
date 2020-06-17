package crmlog

type sourceSystems string

func (sourceSystem sourceSystems) isSystem() sourceSystems {
	return sourceSystem
}

type LogSystem interface {
	isSystem() sourceSystems
}

const (
	CRM_INBOUND   = sourceSystems("CRM_INBOUND")
	CRM_OUTBOUND  = sourceSystems("CRM_OUTBOUND")
	CRM_VALIDATION = sourceSystems("CRM_VALIDATION")
	CRM_INTEGRATION = sourceSystems("CRM_INTEGRATION")
	CRM_EAI   = sourceSystems("CRM_EAI")
	CRM_DATABASE = sourceSystems("CRM_DATABASE")
	CCBS   = sourceSystems("CCBS")
	OMX_MF  = sourceSystems("OMX_MF")
	OMX_CRM = sourceSystems("OMX_CRM")
	OMX_CCBS = sourceSystems("OMX_CCBS")
	CCB_INT   = sourceSystems("CCB_INT")
	VCARE = sourceSystems("VCARE")
	NAS   = sourceSystems("NAS")
	CVSS  = sourceSystems("CVSS")
	IVR = sourceSystems("IVR")
	MVP = sourceSystems("MVP")
	TVS   = sourceSystems("TVS")
	USSD_GW = sourceSystems("USSD_GW")
	ATP = sourceSystems("ATP")
	ESB   = sourceSystems("ESB")
	SBM = sourceSystems("SBM")
)

