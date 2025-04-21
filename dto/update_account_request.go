package dto

type UpdateAccountRequest struct {
	ID          int64  `json:"id"`
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
