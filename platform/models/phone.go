package models

type PhoneInput struct {
	ID       int    `gorm:"id"`
	Phone    string `gorm:"phone"`
	Provider string `gorm:"provider" `
}

func (PhoneInput) TableName() string {
	return "phone_input"
}
