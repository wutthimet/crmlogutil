package crmlog

type productNames string

func (productName productNames) isProduct() productNames {
	return productName
}

type LogProductName interface {
	isProduct() productNames
}

const (
	TrueOnline  = productNames("TrueOnline")
	TrueMoveH   = productNames("TrueMoveH")
	TrueVisions = productNames("TrueVisions")
	IoT         = productNames("IoT")
	Convergence = productNames("Convergence")
	All         = productNames("All")
)
