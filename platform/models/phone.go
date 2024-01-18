package models

type PhoneInput struct {
	ID       int    `gorm:"id" json:"id"`
	Phone    string `gorm:"phone" json:"phone"`
	Provider string `gorm:"provider" json:"provider" `
	IsOdd    int    `gorm:"is_odd" json:"is_odd" `
}

func (PhoneInput) TableName() string {
	return "phone_input"
}
