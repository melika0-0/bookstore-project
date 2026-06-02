package models

type UserRole struct {
	BaseModel 
	User User `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate: NO ACTION;OnDelete:NO ACTION"`
	Role  Role `gorm:"foreignKey:RoleId;references:ID;constraint:OnUpdate: NO ACTION;OnDelete:NO ACTION"`
	UserId int 
    RoleId int 

}