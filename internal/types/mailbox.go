package types

import (
	"log"

	"gorm.io/gorm"
)

type Mailbox struct {
	ID   uint64 `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"size:255"`

	DomainID uint `json:"Domain_id`
	Domain   Domain
}

// overwrite table name
func (Mailbox) TableName() string {
	return "mailbox"
}

func (m *Mailbox) AfterCreate(tx *gorm.DB) (err error) {
	log.Printf("\n-------> Increase mailbox count for %#v", m)
	if err := tx.Model(&Domain{}).Where("id=?", m.DomainID).UpdateColumn("mailbox_count", gorm.Expr("mailbox_count + ?", 1)).Error; err != nil {
		return tx.AddError(err)
	}
	return nil
}
