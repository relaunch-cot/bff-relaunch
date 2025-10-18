package project

import "time"

type CreateProjectPOST struct {
	FreelancerId            string    `json:"freelancerId"`
	Category                string    `json:"category"`
	Name                    string    `json:"name"`
	Description             string    `json:"description"`
	ProjectDeliveryDeadline time.Time `json:"projectDeliveryDeadline"`
	Amount                  float32   `json:"amount"`
}
