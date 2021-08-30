package models

type Reimbursment struct {
	RequestId   uint   `json:"requestId"`
	UserID      uint   `json:"userID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      string `json:"email"`
}
