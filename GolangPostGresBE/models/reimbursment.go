package models

type Reimbursment struct {
	RequestId     uint   `gorm:"primaryKey"`
	UserID        string `json:"userID"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Amount        string `json:"amount"`
	AppovedStatus string `json:"approvalStatus"`
	DateApproved  string `json:"dateApporoved"`
	ApprovedBy    string `json:"approvedBy"`
}
