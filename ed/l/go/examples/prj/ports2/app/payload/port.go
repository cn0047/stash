package payload

// Port represents port entity/DTO (@TBD).
type Port struct {
	ID          string    `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"alpha"`
	City        string    `json:"city" validate:"alpha"`
	Country     string    `json:"country" validate:"alpha"`
	Alias       []string  `json:"alias" validate:"alphanum"`
	Regions     []string  `json:"regions" validate:"alpha"`
	Coordinates []float32 `json:"coordinates" validate:"required"`
	Province    string    `json:"province" validate:"alpha"`
	Timezone    string    `json:"timezone" validate:"required"`
	Unlocs      []string  `json:"unlocs" validate:"required"`
	Code        int       `json:"code" validate:"required"`
}
