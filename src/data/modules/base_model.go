package models

import (
	"time"
	"gorm.io/gorm"
	"database/sql"
)

//for all the models
type BaseModel struct {
	ID        uint `gorm:"primaryKey" json:"id"`

	CreatedAt time.Time `gorm:"Timestamp with timezone; not null" json:"created_at"`
	ModifiedAt sql.NullTime `gorm:"Timestamp with timezone; null" json:"modified_at"`
	DeletedAt sql.NullTime `gorm:"Timestamp with timezone ; null" json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CreatedBy int `gorm:"not null" json:"created_by"`
	ModifiedBy sql.NullInt64 `gorm:"null" json:"modified_by"`
	DeletedBy sql.NullInt64 `gorm:"null" json:"deleted_by"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	Value := tx.Statement.Context.Value("UserId")
	var userId = -1
	if Value != nil {
		userId = Value.(int) 
	}
	// TODO: check user type
	 //value for stuff for logging user
	 m.CreatedAt = time.Now() .UTC()
	m.CreatedBy = userId
	return

}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	Value := tx.Statement.Context.Value("UserId")
	var userId = -1
	if Value != nil {
		userId = Value.(int) 
	}
	
	 m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.ModifiedBy = sql.NullInt64{Int64: int64(userId), Valid: true}
	return

}