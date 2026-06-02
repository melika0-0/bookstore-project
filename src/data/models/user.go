package models

type User struct {
	BaseModel
	Username string `gorm:"type:string;size:20;not null;unique"` //we need thisss
	FirstName string `gorm:"type:string;size:15;null"`
	Lastname string `gorm:"type:string;size:25;null"`
	Email string `gorm:"type:string;size:64;default:null;unique"`
	Password string `gorm:"type:string;size:64;not null"`
	MobileNumber string `gorm:"type:string;size:15;default:null"`
	Enable bool `gorm:"type:boolean;default:true"`
	UserRoles *[]UserRole `gorm:"foreignKey:UserId"`
}