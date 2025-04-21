package dto

type SignUpResponse struct {
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
