package responses

type UserResponse struct{
	ID *int `json:"id"`
	Email *string `json:"email"`
	Name *string `json:"name"`
	Address *string `json:"address"`
}
