package project

import "time"

type CreateProjectPOST struct {
	DeveloperId             string    `json:"developerId"`
	Category                string    `json:"category"`
	ProjectDeliveryDeadline time.Time `json:"projectDeliveryDeadline"`
	Amount                  float32   `json:"amount"`
}
