package logger


type CallerInfoBean struct {
	ClassName string `json:"className"`
	MethodName string `json:"methodName"`
	FileName string `json:"fileName"`
}