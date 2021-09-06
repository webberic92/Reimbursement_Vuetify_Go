package models

type Reimbursment struct {
	RequestId     uint   `json:"requestId"`
	UserID        string `json:"userID"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Amount        string `json:"email"`
	AppovedStatus string `json:"approvalStatus"`
	DateApproved  string `json:"dateApporoved"`
	ApprovedBy    string `json:"approvedBy"`
}
