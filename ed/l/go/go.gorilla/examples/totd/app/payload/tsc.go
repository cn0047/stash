package payload

type defaultInput struct {
	Retailer string
	Env      string
	Token    string
	MFC      string
}

// GetConfigInput holds input parameters to get config form TSC.
type GetConfigInput struct {
	Retailer string // ClientID
	Env      string
	Token    string
	MFC      string
}

// GetFlowRackInput holds input parameters to get FlowRack config.
type IsFlowRacksEnabledInput struct {
	defaultInput
}

// GetFlowRAckByRampID holds input parameters to get FlowRack by specified RampID.
type GetFlowRackByRampIDInput struct {
	defaultInput
	RampID string // e.g. DISPATCH01
}
