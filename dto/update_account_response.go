package dto

type UpdateAccountResponse struct {
	ID          int64  `json:"id"`
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
