package types

type Domain struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Domain       string `json:"domain" gorm:"size:255;unique;not null"`
	MailboxCount int    `json:"mailbox_count" gorm:"not null; default:0"`
}

// overwrite table name
func (Domain) TableName() string {
	return "domain"
}
