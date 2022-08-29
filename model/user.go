package model

type User struct {
	ID        string `json:"id" gorm:"primary_key;"`
	Username  string `json:"username" gorm:"not null;unique;type:varchar(255);"`
	Password  string `json:"-" gorm:"not null;type:varchar(64);"`
	Email     string `json:"email" gorm:"unique;type:varchar(255);default:null;"`
	Phone     string `json:"phone" gorm:"unique;type:varchar(255);default:null;"`
	Role      int    `json:"role" gorm:"type:tinyint;default:1;"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:nano;"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:nano;"`
}

func (u User) TableName() string {
	return "users"
}
