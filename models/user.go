package models

// User model struct
type User struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname" gorm:"type: varchar(255)"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
	Password  string `json:"password" gorm:"type: varchar(255)"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: text"`
	Subscribe string `json:"subscribe" gorm:"type:text"`
	Transaction	[]TransactionInUser `json:"transactions"`
	Role      string `json:"role"`
}

type UsersProfileResponse struct {
	ID			int		`json:"id"`
	Fullname	string	`json:"fullname"`
	Email		string	`json:"email"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Subscribe string `json:"subscribe"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}