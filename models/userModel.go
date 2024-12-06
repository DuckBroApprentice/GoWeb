package models

type User struct {
	Id           int64   `xorm:"pk autoincr" json:"id"`
	UserName string `xorm:"varchar(20)" json:"username"`
	Password string `xorm:"varchar(255)" json:"password"`
	Email    string `xorm:"varchar(255)" json:"email"`
	Phone    string `xorm:"varchar(10)" json:"phone"`
	RegisterTime int64   `xorm:bigint" json:"register_time"`
}
