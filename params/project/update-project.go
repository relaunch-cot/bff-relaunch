package project

import "time"

type UpdateProjectPUT struct {
	UserId                  string    `json:"userId"`
	Name                    string    `json:"name"`
	Description             string    `json:"description"`
	ProjectDeliveryDeadline time.Time `json:"projectDeliveryDeadline"`
	Category                string    `json:"category"`
	Amount                  float32   `json:"amount"`
	UrlImageProject         string    `json:"urlImageProject"`
	Status                  string    `json:"status"`
}
