package model

type Users struct {
	Id         int64  `gorm:"PrimaryKey" json:"id"`
	Username   string `gorm:"type:varchar(255)" json:"username"`
	Password   string `gorm:"type:varchar(255)" json:"password"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Role       string `gorm:"type:text" json:"role"`
	CreatedAt  string `gorm:"type:varchar(100)" json:"created_at"`
	ChangeDate string `gorm:"type:varchar(100)" json:"change_date"`
}

type User struct {
	Id         int64  `gorm:"PrimaryKey" json:"id"`
	Username   string `gorm:"type:varchar(255)" json:"username"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Role       string `gorm:"type:text" json:"role"`
	CreatedAt  string `gorm:"type:varchar(100)" json:"created_at"`
	ChangeDate string `gorm:"type:varchar(100)" json:"change_date"`
}

type ErrorStatus struct {
	StatusCode int64  `json:"status_code"`
	Message    string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
