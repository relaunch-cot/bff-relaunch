package models

type ReportData struct {
	Title    string     `json:"title"`
	Subtitle string     `json:"subtitle,omitempty"`
	Headers  []string   `json:"headers"`
	Rows     [][]string `json:"rows"`
	Footer   string     `json:"footer,omitempty"`
}
