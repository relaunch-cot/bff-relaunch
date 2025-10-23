package project

type RemoveFreelancerFromProjectPATCH struct {
	FreelancerId string `json:"freelancerId"`
	UserId       string `json:"userId"`
}
