package user

type GenerateReportPOST struct {
	Title    string     `json:"title" binding:"required"`
	Subtitle string     `json:"subtitle,omitempty"`
	Headers  []string   `json:"headers" binding:"required"`
	Rows     [][]string `json:"rows"`
	Footer   string     `json:"footer,omitempty"`
}
