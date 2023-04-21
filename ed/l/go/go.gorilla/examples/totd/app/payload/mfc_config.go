package payload

import (
	"fmt"
)

// GetConfigsInput holds input parameters to get MFCConfigs.
type GetConfigsInput struct {
	ClientID string `json:"clientId"`
	Env      string `json:"env"`
	MfcID    string `json:"mfcId"`
}

// GetConfigsOutput holds result for get MFCConfigs.
type GetConfigsOutput struct {
	Configs []MFCConfig `json:"configs"`
}

// MFCConfig represents config for MFC (metadata + configs).
type MFCConfig struct {
	// Metadata fields.
	ClientID  string `json:"client_id" mapstructure:"client_id"`
	Env       string `json:"env" mapstructure:"env"`
	MfcID     string `json:"mfc_id" mapstructure:"mfc_id"`
	UpdatedAt int64  `json:"updated_at" mapstructure:"updated_at"`

	// Config fields.
	// ErrorRamp contains error ramp ID.
	ErrorRamp int64 `json:"error_ramp" mapstructure:"error_ramp"`
	// Count contains value for total available lanes count.
	Count int64 `json:"count" mapstructure:"count"`
	// Depth contains value for lane deep length.
	Depth int64 `json:"depth" mapstructure:"depth"`
	// Start contains value for lane ID to start with.
	Start int64 `json:"start" mapstructure:"start"`
	// IDGen contains formatting string to generate lane ID, like: DISPATCH%02d.
	IDGen string `json:"id_gen" mapstructure:"id_gen"`
	// LaneMapping contains lanes mapping,
	// where key is priority ID and value is actual lane ID.
	// 1 - is the highest priority, which means lane with this priority will be used first.
	// Use LaneMapping to configure arbitrary mapping.
	LaneMapping map[int64]int64 `json:"lane_mapping" mapstructure:"lane_mapping"`
	// ExpressLaneMapping contains express lanes mapping,
	// where key is priority ID and value is actual express lane ID.
	// 1 - is the highest priority, which means lane with this priority will be used first.
	// Use ExpressLaneMapping to configure arbitrary mapping for express lanes.
	ExpressLaneMapping map[int64]int64 `json:"express_lane_mapping" mapstructure:"express_lane_mapping"`
	// FlowRacksMapping contains flow racks mapping.
	FlowRacksMapping map[int64]string `json:"flow_racks_mapping" mapstructure:"flow_racks_mapping"`
}

// GetCount gets lanes count value.
func (m *MFCConfig) GetCount() int64 {
	if len(m.LaneMapping) > 0 {
		return int64(len(m.LaneMapping))
	}

	return m.Count
}

// GetExpressCount gets express lanes count value.
func (m *MFCConfig) GetExpressCount() int64 {
	if len(m.ExpressLaneMapping) > 0 {
		return int64(len(m.ExpressLaneMapping))
	}

	return 0
}

// GetRampStringID gets ramp string ID, like: DISPATCH01.
func (m *MFCConfig) GetRampStringID(laneID int64) string {
	return fmt.Sprintf(m.IDGen, laneID)
}

// GetErrorRampStringID gets error ramp string ID, like: DISPATCH99.
func (m *MFCConfig) GetErrorRampStringID() string {
	return fmt.Sprintf(m.IDGen, m.ErrorRamp)
}

// ConvertLaneIdxToID converts lane index (internal iterator index) to lane ID (human understandable ID).
func (m *MFCConfig) ConvertLaneIdxToID(idx int64, isExpress bool) (string, error) {
	// Case when provided index is error ramp.
	if idx == m.ErrorRamp {
		return m.GetErrorRampStringID(), nil
	}

	// Case when FlowRacks enabled.
	if len(m.FlowRacksMapping) > 0 {
		laneID, ok := m.FlowRacksMapping[idx]
		if !ok {
			return "", fmt.Errorf("failed to find lane in flowRacks mapping")
		}
		return laneID, nil
	}

	// -1 intentionally here, because 0 can be valid value,
	// so with purpose to avoid possible collisions this value contains definitely incorrect value.
	laneID := int64(-1)

	// Express order.
	if isExpress {
		if len(m.ExpressLaneMapping) == 0 {
			return "", fmt.Errorf("express lanes mapping is empty in MFC config")
		}
		lID, ok := m.ExpressLaneMapping[idx]
		if !ok {
			return "", fmt.Errorf("failed to find lane in express config mapping")
		}

		return m.GetRampStringID(lID), nil
	}

	// Regular order.
	if len(m.LaneMapping) > 0 { // Case when lane order is defined by mapping.
		lID, ok := m.LaneMapping[idx]
		if !ok {
			return "", fmt.Errorf("failed to find lane in config mapping")
		}
		laneID = lID
	} else if m.Start == 1 { // Case when lane order is A-Z.
		laneID = idx
	} else if m.Start == m.Count { // Case when lane order is Z-A.
		laneID = m.Start - idx + 1 // +1 to get 1 not 0.
	}

	if laneID == -1 {
		return "", fmt.Errorf("unsupported case for idx: %v, isExpress: %t", idx, isExpress)
	}

	return m.GetRampStringID(laneID), nil
}
