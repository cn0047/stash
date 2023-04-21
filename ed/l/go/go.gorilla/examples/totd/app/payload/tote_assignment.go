package payload

const (
	// ViewGroupedByOrder contains name value for list view which is grouped by order.
	ViewGroupedByOrder = "groupedByOrder"
)

// ToteAssignment represents ToteAssignment entity.
type ToteAssignment struct {
	ID        string `json:"id" spanner:"id"`
	ClientID  string `json:"clientId" spanner:"client_id"`
	MfcID     string `json:"mfcId" spanner:"mfc_id"`
	OrderID   string `json:"orderId" spanner:"order_id"`
	IsExpress bool   `json:"isExpress" spanner:"is_express"`
	ToteID    string `json:"toteId" spanner:"tote_id"`
	LaneIdx   int64  `json:"laneIdx" spanner:"lane_idx"`
	LaneID    string `json:"laneId" spanner:"lane_id"`
	CreatedAt int64  `json:"createdAt" spanner:"created_at"`
}

// ListToteAssignmentsInput holds input parameters to list ToteAssignments.
type ListToteAssignmentsInput struct {
	ClientID string `json:"clientId"`
	MfcID    string `json:"mfcId"`
	View     string `json:"view"`
	OrderID  string `json:"orderId"`
}

// ListToteAssignmentsOutput holds result for list ToteAssignments.
type ListToteAssignmentsOutput struct {
	ToteAssignments []ToteAssignment `json:"toteAssignments"`
}

// ListTotesGroupedByLaneOutput holds result for list ToteAssignments grouped by lane.
type ListTotesGroupedByLaneOutput struct {
	RegularLanes map[int64][]string `json:"regularLanes"`
	ExpressLanes map[int64][]string `json:"expressLanes"`
}

// ListToteAssignmentsGroupedByOrderOutput holds result for list ToteAssignments grouped by order.
type ListToteAssignmentsGroupedByOrderOutput struct {
	Balance map[string][]string `json:"balance"`
	Orders  map[string][]string `json:"orders"`
}

// CreateToteAssignmentInput holds input parameters to create ToteAssignment.
type CreateToteAssignmentInput struct {
	ID        string `json:"id"`
	ClientID  string `json:"clientId"`
	MfcID     string `json:"mfcId"`
	OrderID   string `json:"orderId"`
	ToteID    string `json:"toteId"`
	IsExpress bool   `json:"isExpress"`
	DryRun    bool   `json:"dryRun"`
}

// CreateToteAssignmentOutput holds output result of ToteAssignment creation.
type CreateToteAssignmentOutput struct {
	ToteAssignment ToteAssignment `json:"toteAssignment"`
}

// UpdateToteAssignmentInput holds input parameters to update ToteAssignment.
type UpdateToteAssignmentInput struct {
	ID       string `json:"id"`
	ClientID string `json:"clientId"`
	MfcID    string `json:"mfcId"`
	LaneID   string `json:"laneId"`
	LaneIdx  int64  `json:"laneIdx"`
	ToteID   string `json:"toteId"`
}

// UpdateToteAssignmentOutput holds output result for ToteAssignment update.
type UpdateToteAssignmentOutput struct {
	ToteAssignment ToteAssignment `json:"toteAssignment"`
}

// DeleteToteAssignmentInput holds input parameters to delete ToteAssignment.
type DeleteToteAssignmentInput struct {
	ClientID string
	MfcID    string
	ToteIDs  []string
}

type DeleteToteAssignmentOutput struct {
}
