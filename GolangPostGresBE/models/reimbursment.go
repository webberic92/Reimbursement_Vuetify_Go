package models

type Reimbursment struct {
	RequestId      uint64 `gorm:"primaryKey"`
	UserID         uint64 `gorm:"foreignKey:User;references:Id;" json:"userID"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Amount         string `json:"amount"`
	ApprovedStatus string `json:"approvalStatus"`
	DateApproved   string `json:"dateApproved"`
	ApprovedBy     uint64 `gorm:"foreignKey:User;references:Id;" json:"approved_by"`
}
