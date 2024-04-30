package request

type CreateUserRequest struct {
	Name        string `json:"name"`
	Affiliation string `json:"affiliation"`
}
