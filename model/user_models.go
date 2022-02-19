package model

// struct user
type Users struct {
	ID           uint   `json:"id" form:"id"`
	Nama         string `json:"nama" form:"nama"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Phone_Number string `json:"phone" form:"phone"`
	Token        string `json:"token" form:"token"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Get_User struct {
	ID           uint
	Nama         string
	Email        string
	Phone_Number string
	Token        string
}
