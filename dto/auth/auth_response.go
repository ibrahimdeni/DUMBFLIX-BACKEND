package authdto

type LoginResponse struct {
	ID       int    `json:"id"`
	Fullname string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" from:"email"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	Gender  string `json:"gender" gorm:"type: varchar(255)"`
	Phone   string `json:"phone" gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: varchar(255)"`
	Subscribe string `json:"subscribe" form:"subscribe"`
	Role string `json:"role"`
}

type CheckAuthResponse struct {
	ID        int    `json:"id"`
	Fullname  string `gorm:"type: varchar(255)" json:"fullname"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Gender    string `gorm:"type: varchar(255)" json:"gender"`
	Phone     string `gorm:"type: varchar(255)" json:"phone"`
	Address   string `gorm:"type: varchar(255)" json:"address"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
	Subscribe string `json:"subscribe" form:"subscribe"`
	Role      string `gorm:"type: varchar(100)" json:"role"`
}
